# Define the name of the binary
BINARY_NAME := gohooks

# Define the default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	go build -o bin/$(BINARY_NAME)

# Run tests
.PHONY: test
test:
	go test -v ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -f bin/$(BINARY_NAME)
	rm -rf bin/

# Run the application
.PHONY: run
run: build
	./bin/$(BINARY_NAME)

# Format code
.PHONY: fmt
fmt:
	go fmt ./...
	go mod tidy

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Generate documentation
.PHONY: doc
doc:
	godoc -http :8080

# Run security checks
.PHONY: security
security:
	gosec ./...

# Run static analysis
.PHONY: static-analysis
static-analysis: lint vet

# Vet code
.PHONY: vet
vet:
	go vet ./...

# Tidy go.mod and go.sum
.PHONY: tidy
tidy:
	go mod tidy

# Pre commit hook
# todo: need to add security later and fix the issues
.PHONY: pre-commit
pre-commit: fmt lint vet build test

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all                - Default target, builds the application"
	@echo "  build              - Build the application"
	@echo "  test               - Run tests"
	@echo "  clean              - Remove build artifacts"
	@echo "  run                - Build and run the application"
	@echo "  fmt                - Format code and tidy modules"
	@echo "  lint               - Lint code"
	@echo "  doc                - Generate documentation"
	@echo "  static-analysis    - Run static analysis (lint and vet)"
	@echo "  vet                - Vet code"
	@echo "  tidy               - Tidy go.mod and go.sum"
	@echo "  pre-commit         - Run pre-commit checks (fmt, lint, vet, build, test, security)"
	@echo "  help               - Show this help message"
