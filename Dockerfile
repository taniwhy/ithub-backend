FROM golang:1.13

ENV GO111MODULE=on

RUN apt-get -y update
RUN apt-get install -y \
    curl \
    gnupg

RUN curl -sL https://deb.nodesource.com/setup_11.x | bash -
RUN apt-get install -y nodejs
RUN npm install npm@latest -g

RUN npm install -g aglio --unsafe-perm

RUN mkdir -p /go/src/github.com/taniwhy/ithub-backend
WORKDIR /go/src/github.com/taniwhy/ithub-backend

COPY ./ ./

RUN go get github.com/rubenv/sql-migrate/...

EXPOSE 8000

CMD go run ./cmd/app/main.go
