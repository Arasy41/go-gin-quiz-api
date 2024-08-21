# Variables
APP_NAME=go-gin-quiz-api
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")
PORT=8080

# Default target
all: build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod vendor

# Build the application
build: deps
	@echo "Building the application..."
	@go build -o $(APP_NAME) cmd/api/main.go

# Run the application with build
run: build
	@echo "Running the application on port $(PORT)..."
	@./$(APP_NAME)

# Run the application
run-dev:
	@echo "Running the application on port $(PORT)..."
	@go run cmd/api/main.go

# Clean the build
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)

# Tidy a Module
tidy:
	@echo "Tidying the application..."
	@go mod tidy

# Init Swagger
swagger:
	@echo "Generating swagger docs..."
	swag init -g cmd/api/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Run tests with coverage
test-cover:
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

# Run the application in Docker
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):latest .

docker-run:
	@echo "Running Docker container..."
	@docker run -p $(PORT):$(PORT) $(APP_NAME):latest

# Format the code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint the code
lint:
	@echo "Linting code..."
	@golangci-lint run

# Help
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all            Build the application (default)"
	@echo "  deps           Install dependencies"
	@echo "  build          Build the application"
	@echo "  run            Build and run the application"
	@echo "  clean          Clean the build"
	@echo "  test           Run tests"
	@echo "  test-cover     Run tests with coverage"
	@echo "  docker-build   Build Docker image"
	@echo "  docker-run     Run Docker container"
	@echo "  fmt            Format the code"
	@echo "  lint           Lint the code"
	@echo "  help           Show this help message"

.PHONY: all deps build run clean test test-cover docker-build docker-run fmt lint help
