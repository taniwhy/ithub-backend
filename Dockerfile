FROM golang:1.13

ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/taniwhy/ithub-backend
WORKDIR /go/src/github.com/taniwhy/ithub-backend
ADD . /go/src/github.com/taniwhy/ithub-backend

RUN go get github.com/rubenv/sql-migrate/...

EXPOSE 8000

CMD go run main.go

