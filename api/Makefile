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

run-local:
	make docker-deps-start
	PATH=$(PATH):$(HOME)/go/bin godotenv -f .env.local go run .
.PHONY: local

docker-deps-start:
	COMPOSE_PROJECT_NAME=test docker-compose up -d
.PHONY: docker-deps-start

docker-deps-stop:
	COMPOSE_PROJECT_NAME=test docker-compose down
.PHONY: docker-deps-stop

docker-deps-restart:
	make docker-deps-stop
	make docker-deps-start
.PHONY: docker-deps-restart

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