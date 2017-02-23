FROM golang:1.6.2-alpine

MAINTAINER TFG Co <backend@tfgco.com>

EXPOSE 80

RUN apk update
RUN apk add git make g++ apache2-utils
RUN apk add --update bash

RUN go get -u github.com/Masterminds/glide/...
RUN go get -u github.com/topfreegames/goose/cmd/goose

ADD . /go/src/github.com/topfreegames/khan

WORKDIR /go/src/github.com/topfreegames/khan
RUN glide install
RUN go install github.com/topfreegames/khan

ENV KHAN_POSTGRES_HOST 0.0.0.0
ENV KHAN_POSTGRES_PORT 5432
ENV KHAN_POSTGRES_USER khan
ENV KHAN_POSTGRES_PASSWORD ""
ENV KHAN_POSTGRES_DBNAME khan
ENV KHAN_ELASTICSEARCH_HOST 0.0.0.0
ENV KHAN_ELASTICSEARCH_PORT 9200
ENV KHAN_ELASTICSEARCH_INDEX khan
ENV KHAN_ELASTICSEARCH_SNIFF false
ENV KHAN_WEBHOOKS_WORKERS 50
ENV KHAN_WEBHOOKS_RUNSTATS false
ENV KHAN_WEBHOOKS_STATSPORT 9999
ENV KHAN_REDIS_HOST 0.0.0.0
ENV KHAN_REDIS_PORT 6379
ENV KHAN_REDIS_DATABASE 0
ENV KHAN_REDIS_POOL 30
ENV KHAN_REDIS_PASSWORD ""

ENV KHAN_SENTRY_URL ""
ENV KHAN_BASICAUTH_USERNAME ""
ENV KHAN_BASICAUTH_PASSWORD ""

CMD /go/bin/khan start --bind 0.0.0.0 --port 80 --fast --config /go/src/github.com/topfreegames/khan/config/default.yaml
