FROM golang:1.13

ENV DOCKERIZE_VERSION v0.6.1
ENV GO111MODULE=on

RUN apt-get update && apt-get install -y wget \
    && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN mkdir -p /go/src/github.com/taniwhy/ithub-backend
WORKDIR /go/src/github.com/taniwhy/ithub-backend
ADD . /go/src/github.com/taniwhy/ithub-backend

RUN go get github.com/rubenv/sql-migrate/...
RUN sql-migrate up

EXPOSE 8000

CMD go run main.go

