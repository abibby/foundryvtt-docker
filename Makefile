all:
	make build
	make push

build: has-version
	docker build . -t abibby/foundryvtt -t abibby/foundryvtt:$(VERSION)

push: has-version
	docker push abibby/foundryvtt
	docker push abibby/foundryvtt:$(VERSION)

has-version:
	@test -n '$(VERSION)' || (printf "\nMissing argument VERSION\n\n" && exit 1)


run:
	docker run -p 30000:30000 -v $(PWD)/data:/data abibby/foundryvtt