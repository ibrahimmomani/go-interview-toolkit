# Golang Interview Kit

A comprehensive, production-quality collection of data structures and algorithms for technical interviews, implemented in Go with full generic type support.

## 🚀 Features

- **Generic Type Support** - Works with any type using Go 1.21+ generics
- **Memory Efficient** - Optimized implementations with proper capacity management
- **Production Ready** - Comprehensive error handling and edge case coverage
- **Interview Focused** - Includes common algorithmic patterns and use cases
- **100% Test Coverage** - Extensive test suites with benchmarks
- **Documentation** - Clear API documentation with time complexity notes

## 📦 Data Structures Implemented

### Linear Collections (`collections/`)

#### 1. LinkedList
- **Operations**: Append, Prepend, Insert, Delete, Get, Find
- **Features**: Head/tail pointers for O(1) operations, reverse capability
- **Use Cases**: Dynamic insertion/deletion, when array resizing is expensive

#### 2. Stack (Coming soon)
- **Operations**: Push, Pop, Peek, MultiPush, MultiPop
- **Features**: LIFO behavior, dynamic resizing, capacity optimization
- **Use Cases**: DFS, expression evaluation, backtracking, balanced parentheses

#### 3. Queue (Coming soon)
- **Operations**: Enqueue, Dequeue, Front, Rear, MultiEnqueue, MultiDequeue
- **Features**: FIFO behavior, circular buffer implementation, auto-resizing  
- **Use Cases**: BFS, level-order traversal, sliding window, process scheduling

#### 4. Deque (Double-ended Queue) (Coming soon)
- **Operations**: PushFront, PushBack, PopFront, PopBack, Rotate
- **Features**: O(1) operations at both ends, can work as stack or queue
- **Use Cases**: Sliding window maximum, palindrome checking, undo/redo systems

## 🛠 Installation & Usage

### Setup
```bash
git clone https://github.com/ibrahimmomani/go-interview-toolkit
cd go-interview-toolkit
go mod tidy
```

### Basic Usage
```go
import "github.com/ibrahimmomani/go-interview-toolkit/collections"

func main() {
    // LinkedList
    list := collections.NewLinkedList[int]()
    list.Append(1)
    list.Append(2)
}
```

## 🧪 Testing & Building

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run benchmarks
make bench

# Build examples
make build

# Run example demos
make run-examples

# Format code
make fmt

# Clean build artifacts
make clean
```

## 📊 Performance Characteristics

| Data Structure | Access | Search | Insertion | Deletion | Space |
|---------------|--------|--------|-----------|----------|-------|
| LinkedList    | O(n)   | O(n)   | O(1)*     | O(1)*    | O(n)  |

*\* At head/tail for LinkedList*

## 🎯 Interview Problem Examples

Each data structure includes demo code showing how to use.

## 🏗 Project Structure

```
go-interview-toolkit/
├── collections/              # Linear data structures
│   ├── linkedlist.go        # Singly linked list
│   └── *_test.go            # Comprehensive tests
├── examples/                # Usage demonstrations
│   ├── linkedlist_demo.go
├── Makefile                 # Build automation
└── README.md
```


## 📋 Best Practices Implemented

- **Go Conventions**: Proper naming, error handling, zero values
- **Memory Management**: Efficient allocation, GC-friendly patterns
- **Type Safety**: Full generic support with proper constraints
- **Testing**: Table-driven tests, benchmarks, edge cases
- **Documentation**: Godoc comments with complexity analysis
- **CI/CD Ready**: Makefile automation, formatting, linting

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Add comprehensive tests
4. Follow Go best practices
5. Submit a pull request

## 📖 Documentation

Each package includes:
- Godoc-compatible documentation
- Time/space complexity notes
- Usage examples
- Common pitfalls and solutions

## 📝 License

MIT License - see LICENSE file for details.

## 🎯 Why This Toolkit?

**Interview Ready**: Covers 90% of data structure questions in technical interviews

**Production Quality**: Not just working code, but industry-standard implementations

**Educational**: Learn Go best practices while mastering algorithms

**Comprehensive**: From basic operations to advanced algorithmic patterns

**Performant**: Optimized implementations with proper complexity characteristics

---

*Built with ❤️ for the Go community and interview preparation*