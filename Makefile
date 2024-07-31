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
	go test ./...

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

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Generate documentation
.PHONY: doc
doc:
	godoc -http :8080

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  test     - Run tests"
	@echo "  clean    - Remove build artifacts"
	@echo "  run      - Build and run the application"
	@echo "  fmt      - Format code"
	@echo "  lint     - Lint code"
	@echo "  doc      - Generate documentation"
	@echo "  help     - Show this help message"
