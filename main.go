package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if errors.Is(err, os.ErrNotExist) {
	} else if err != nil {
		log.Fatal(err)
	}

	buildFolder := "./builds"

	build, err := DownloadBuild(buildFolder)
	if err != nil {
		log.Printf("failed to download new build: %v", err)

		builds, err := os.ReadDir(buildFolder)
		if err != nil {
			log.Fatal(err)
		}
		for _, b := range builds {
			if b.Name() > build {
				build = b.Name()
			}
		}
	}

	for {

		ctx, cancel := context.WithCancel(context.Background())
		if os.Getenv("FOUNDRY_BUILD") == "" {
			go func() {
				t := time.Now()
				date := time.Date(t.Year(), t.Month(), t.Day(), 16, 0, 0, 0, t.Location())
				delay := (date.Sub(t) + time.Hour*24) % (time.Hour * 24)
				time.Sleep(delay)

				newBuild, err := DownloadBuild(buildFolder)
				if err != nil {
					log.Printf("failed to download new build: %v", err)
					return
				}
				if newBuild == build {
					return
				}
				build = newBuild
				cancel()
			}()
		}

		cmd := exec.CommandContext(ctx, "node", path.Join(buildFolder, build, "resources/app/main.js"), "--dataPath=./data2")

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
