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
