FROM alpine:3.12.0 AS unzip

COPY foundryvtt-0.6.4.zip foundryvtt.zip

RUN mkdir /foundryvtt && unzip foundryvtt.zip -d /foundryvtt

FROM node:14.4.0-alpine3.12

COPY --from=unzip /foundryvtt/resources/app /foundryvtt

EXPOSE 30000

VOLUME /data /config

CMD [ "node", "/foundryvtt/main.js", "--dataPath=/data" ]
