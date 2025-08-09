# Makefile
.PHONY: test build clean fmt vet coverage help run-examples

# Default target
all: fmt vet test

# Run tests (exclude examples folder)
test:
	go test -v ./collections/... 

# Run tests with coverage (exclude examples folder)
coverage:
	go test -v -coverprofile=coverage.out ./collections/... 
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build examples (each example is in its own folder)
build:
	mkdir -p bin
	go build -o bin/linkedlist_demo ./examples/linkedlist/
	go build -o bin/stack_demo ./examples/stack/

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Run benchmarks
bench:
	go test -bench=. -benchmem ./collections/...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run examples
run-examples: build
	@echo "Running LinkedList demo:"
	./bin/linkedlist_demo
	@echo "\n===========================================\n"
	@echo "Running Stack demo:"
	./bin/stack_demo

# Run individual examples
run-linkedlist: build
	./bin/linkedlist_demo

run-stack: build
	./bin/stack_demo

# Check if examples compile without building
check-examples:
	@echo "Checking if all examples compile..."
	go build -o /dev/null ./examples/linkedlist/
	go build -o /dev/null ./examples/stack/
	@echo "All examples compile successfully!"

# Help
help:
	@echo "Available targets:"
	@echo "  test                    - Run all tests"
	@echo "  coverage                - Run tests with coverage report"
	@echo "  build                   - Build all example applications"
	@echo "  run-examples           - Build and run all example applications"
	@echo "  run-linkedlist         - Run LinkedList demo only"
	@echo "  run-stack              - Run Stack demo only"
	@echo "  check-examples         - Check if examples compile"
	@echo "  fmt                    - Format code"
	@echo "  vet                    - Run go vet"
	@echo "  clean                  - Clean build artifacts"
	@echo "  bench                  - Run benchmarks"
	@echo "  deps                   - Install and tidy dependencies"
	@echo "  all                    - Run fmt, vet, and test"