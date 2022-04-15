.PHONY: run lint build
run:
	go run main.go

lint:
	golangci-lint run

build:
	go build
