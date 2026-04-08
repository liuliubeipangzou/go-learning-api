APP_NAME=go-learning-api

.PHONY: run test build fmt tidy

run:
	go run ./cmd/server

test:
	go test ./...

build:
	mkdir -p bin
	go build -o ./bin/$(APP_NAME) ./cmd/server

fmt:
	gofmt -w ./cmd ./internal

tidy:
	go mod tidy
