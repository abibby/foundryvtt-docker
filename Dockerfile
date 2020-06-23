FROM alpine:3.12.0 AS unzip

COPY foundryvtt-0.6.2.zip foundryvtt.zip

RUN mkdir /foundryvtt && unzip foundryvtt.zip -d /foundryvtt

FROM node:14.4.0-alpine3.12

COPY --from=unzip /foundryvtt/resources/app /foundryvtt
# COPY foundryvtt/resources/app /foundryvtt

COPY entry.sh /entry.sh
COPY entry.js /entry.js

VOLUME /data /config

ENTRYPOINT [ "/entry.sh" ]

CMD [ "node", "/foundryvtt/main.js", "--dataPath=/data" ]
