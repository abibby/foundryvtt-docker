FROM golang:1.20 AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /foundryvtt-runner ./...

FROM node:19-bullseye

COPY --from=build /foundryvtt-runner /usr/local/bin/foundryvtt-runner

EXPOSE 30000

VOLUME /data /config /builds

ENTRYPOINT ["/bin/sh", "-c"]

CMD [ "/usr/local/bin/foundryvtt-runner" ]
