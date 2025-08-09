# Makefile
.PHONY: test build clean fmt vet coverage help run-examples

# Default target
all: fmt vet test

# Run tests (only LinkedList and Stack)
test:
	go test -v ./collections/...

# Run tests with coverage (only LinkedList and Stack)
coverage:
	go test -v -coverprofile=coverage.out ./collections/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build examples (only LinkedList and Stack)
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

# Run benchmarks (only LinkedList and Stack)
bench:
	go test -bench=BenchmarkAppend -bench=BenchmarkPrepend -bench=BenchmarkPush -bench=BenchmarkPop -benchmem ./collections

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run examples (only LinkedList and Stack)
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

# Check if examples compile without building (only LinkedList and Stack)
check-examples:
	@echo "Checking if LinkedList and Stack examples compile..."
	go build -o /dev/null ./examples/linkedlist/
	go build -o /dev/null ./examples/stack/
	@echo "LinkedList and Stack examples compile successfully!"

# Test specific data structures
test-linkedlist:
	go test -v ./collections -run "TestLinkedList"

test-stack:
	go test -v ./collections -run "TestStack"

# Benchmark specific data structures
bench-linkedlist:
	go test -bench=BenchmarkAppend -bench=BenchmarkPrepend -bench=BenchmarkGet -bench=BenchmarkFind -benchmem ./collections

bench-stack:
	go test -bench=BenchmarkPush -bench=BenchmarkPop -bench=BenchmarkPeek -benchmem ./collections

# Help
help:
	@echo "Available targets (LinkedList and Stack only):"
	@echo "  test                  - Run LinkedList and Stack tests"
	@echo "  test-linkedlist       - Run LinkedList tests only"
	@echo "  test-stack            - Run Stack tests only"
	@echo "  coverage              - Run tests with coverage report"
	@echo "  build                 - Build LinkedList and Stack examples"
	@echo "  run-examples          - Run LinkedList and Stack demos"
	@echo "  run-linkedlist        - Run LinkedList demo only"
	@echo "  run-stack             - Run Stack demo only"
	@echo "  bench                 - Run basic benchmarks for both"
	@echo "  bench-linkedlist      - Run LinkedList benchmarks"
	@echo "  bench-stack           - Run Stack benchmarks"
	@echo "  check-examples        - Check if examples compile"
	@echo "  fmt                   - Format code"
	@echo "  vet                   - Run go vet"
	@echo "  clean                 - Clean build artifacts"
	@echo "  deps                  - Install and tidy dependencies"
	@echo "  all                   - Run fmt, vet, and test"