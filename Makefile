all:

build:
	docker build . -t abibby/foundryvtt

run:
	docker run -p 30000:30000 -v $(PWD)/data:/data abibby/foundryvtt