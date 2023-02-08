package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func DownloadBuild(buildFolder string) (string, error) {

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Jar: jar,
	}

	err = login(client)
	if err != nil {
		return "", err
	}

	build := os.Getenv("FOUNDRY_BUILD")
	if build == "" {
		build, err = getBuild(client)
		if err != nil {
			return "", err
		}
	}

	activeBuild := path.Join(buildFolder, build)

	if folderExists(activeBuild) {
		return build, nil
	}

	resp, err := client.Get(fmt.Sprintf("https://foundryvtt.com/releases/download?build=%s&platform=linux", build))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	zipFile := "./foundryvtt.zip"

	f, err := os.Create(zipFile)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(buildFolder, os.ModePerm)
	if err != nil {
		return "", err
	}
	err = Extract(zipFile, activeBuild)
	if err != nil {
		return "", err
	}

	return build, nil
}

func login(client *http.Client) error {

	resp, err := client.Get("https://foundryvtt.com/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("failed to login")
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return err
	}
	csrf, ok := doc.Find(`#login-form [name="csrfmiddlewaretoken"]`).Attr("value")
	if !ok {
		return errors.New("Failed to find csrf")
	}

	resp, err = postForm(client, "https://foundryvtt.com/auth/login/", url.Values{
		"csrfmiddlewaretoken": {csrf},
		"login_redirect":      {"/"},
		"login_username":      {os.Getenv("FOUNDRY_EMAIL")},
		"login_password":      {os.Getenv("FOUNDRY_PASSWORD")},
		"login":               {""},
	})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("failed to login")
	}

	return nil
}

func getBuild(client *http.Client) (string, error) {
	resp, err := client.Get(fmt.Sprintf("https://foundryvtt.com/community/%s/licenses", os.Getenv("FOUNDRY_USERNAME")))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "", err
	}
	build, ok := doc.Find(`select[name="build"] option`).Attr("value")
	if !ok {
		fmt.Print(doc.Html())
		return "", errors.New("Failed to find build")
	}

	return build, nil
}

func postForm(c *http.Client, url string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://foundryvtt.com/")
	return c.Do(req)
}

func folderExists(name string) bool {
	_, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}
