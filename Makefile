# Conditional variables, can be overridden at runtime
DOCKER_COMPOSE_FILE ?= docker-compose.yml
GO_BINARY ?= go
VERSION := "development" # Hardcoded Version
VCS_URL := "https://github.com/ashwin2125/go-mathify" # Hardcoded VCS_URL
BUILD_DATE := $(shell date '+%a|%d/%b/%Y|%H:%M:%S')

# Phony targets don't correspond to filenames
.PHONY: build test docker-build docker-push clean dep up down logs

all: dep build test docker-build docker-push up logs clean

# Build the Go-Mathify binary
build:
	@echo "Building the Go-Mathify binary"
	@$(GO_BINARY) build -o go-mathify ./cmd/go-mathify

# Run all tests
test:
	@echo "Running all tests"
	@$(GO_BINARY) test -v -parallel=4 ./pkg/...

# Build the Docker image
docker-build:
	@docker-compose build --build-arg VERSION=$(VERSION) --build-arg BUILD_DATE="$(BUILD_DATE)" --build-arg VCS_URL=$(VCS_URL)
	@echo "Build successful!"
	@echo "Version: $(VERSION)"
	@echo "Build Date: $(BUILD_DATE)"
	@echo "VCS URL: $(VCS_URL)"

# Push the Docker image to a registry
docker-push:
	@echo "Pushing Docker image to registry"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) push

# Clean up generated files
clean:
	@echo "Cleaning up generated files"
	@rm -f go-mathify

# Sort out Go mod dependencies
dep:
	@echo "Sorting out Go module dependencies"
	@$(GO_BINARY) mod tidy && $(GO_BINARY) mod download && $(GO_BINARY) mod verify

# Run the app using Docker Compose
up:
	@echo "Starting up services with Docker Compose"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop the Docker Compose services
down:
	@echo "Stopping services with Docker Compose"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Tail application logs
logs:
	@echo "Following logs for go-mathify service"
	@docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f go-mathify
