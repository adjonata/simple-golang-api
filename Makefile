include ./.env

.PHONY: swagger
swagger:
	./scripts/swagger.sh

.PHONY: clean-db
clean-db:
	$(SUDO) rm -rf ./test.db

# Deploy, pipes and tests
GO=go
DOCKER=docker
DOCKER_COMPOSE=docker-compose
DOCKER_IMAGE=api-test

.PHONY: test
test:
	$(GO) test ./...

.PHONY: run
run:
	$(GO) run main.go

.PHONY: build
build:
	$(GO) build -o build/main main.go

.PHONY: docker-build
docker-build:
	$(DOCKER) build --rm --tag $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	$(DOCKER) run --env-file .env $(DOCKER_IMAGE) .

.PHONY: start-docker-compose
start-docker-compose:
	$(DOCKER_COMPOSE) up -d

.PHONY: stop-docker-compose
stop-docker-compose:
	$(DOCKER_COMPOSE) down

.PHONY: all
all: clean-db build run
