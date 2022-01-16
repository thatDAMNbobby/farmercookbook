get:
	go get -t -v
	go get github.com/joho/godotenv/cmd/godotenv # This lets you run godotenv from the cli
	go get github.com/golangci/golangci-lint/cmd/golangci-lint # Run lint
.PHONY: get

mocks:
	rm -rf generatedmocks/*
.PHONY: mocks

build:
	go build -v ./...
.PHONY: build

lint:
	PATH=$(PATH):$(HOME)/go/bin golangci-lint run
.PHONY: lint

test:
	go test -coverprofile=cover.out `go list ./... | grep -v generatedmocks`
.PHONY: test

fmt:
	gofmt -s -w .
.PHONY: fmt

precommit:
	make mocks
	make test
	make lint
	make fmt
.PHONY: precommit

shortprecommit:
	make test
	make lint
	make fmt
.PHONY: shortprecommit