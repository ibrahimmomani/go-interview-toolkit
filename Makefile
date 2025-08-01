# Makefile
.PHONY: test build clean fmt vet lint coverage help

# Default target
all: fmt vet test

# Run tests
test:
	go test -v ./...

# Run tests with coverage
coverage:
	go test -v -coverprofile=coverage.out ./collections/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build examples
build:
	mkdir -p bin
	go build -o bin/linkedlist_demo ./examples/linkedlist_demo.go

# Run examples
run-examples: build
	@echo "Running LinkedList demo:"
	./bin/linkedlist_demo

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run golint (requires golint to be installed)
lint:
	golint ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Run benchmarks
bench:
	go test -bench=. -benchmem ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Help
help:
	@echo "Available targets:"
	@echo "  test         - Run all tests"
	@echo "  coverage     - Run tests with coverage report"
	@echo "  build        - Build example applications"
	@echo "  run-examples - Build and run example applications"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  lint         - Run golint"
	@echo "  clean        - Clean build artifacts"
	@echo "  bench        - Run benchmarks"
	@echo "  deps         - Install and tidy dependencies"
	@echo "  all          - Run fmt, vet, and test"