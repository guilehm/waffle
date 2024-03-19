FROM golang:1.22.1-alpine AS builder

RUN set -ex && apk add --no-progress --no-cache \
      gcc \
      musl-dev

ADD . /go/src
WORKDIR /go/src/tmdb

RUN  go mod download
COPY . .

RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -a -v -tags musl
