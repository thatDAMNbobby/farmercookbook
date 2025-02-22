APP_NAME ?= app

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	go test -race -v -timeout 30s ./...

.PHONY: dev
dev:
	go build -o ./tmp/main ./cmd/main.go && air

.PHONY: build
build:
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/main.go

.PHONY: docker-build
docker-build:
	docker-compose -f ./dev/docker-compose.yml build

.PHONY: docker-up
docker-up:
	docker-compose -f ./dev/docker-compose.yml up

.PHONY: docker-dev
docker-dev:
	docker-compose -f ./dev/docker-compose.yml -f ./dev/docker-compose.dev.yml up

.PHONY: docker-down
docker-down:
	docker-compose -f ./dev/docker-compose.yml down

.PHONY: docker-clean
docker-clean:
	docker-compose -f ./dev/docker-compose.yml down -v --rmi all