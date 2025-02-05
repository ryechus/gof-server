SERVICE_NAME?=gof-server

.PHONY: all
all: test

.PHONY: test
test: test-integration

.PHONY: build-docker
build-docker:
	docker build -t ${SERVICE_NAME}:latest --build-arg SERVICE_NAME="${SERVICE_NAME}" -f ./deployments/docker/Dockerfile .

.PHONY: test-integration
test-integration: build-docker
	go test -v -timeout=10s ./...
