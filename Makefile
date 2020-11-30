# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

run:
	go run cmd/app/main.go
test:
	$(GOTEST) -v ./...

aglio:
	aglio -i ./api/app.apib -o ./web/html/api.html -t "slate"