FROM golang:1.22.1-alpine

ADD . /go/src
WORKDIR /go/src

RUN go get -d -v ./...
RUN go install -v ./...
