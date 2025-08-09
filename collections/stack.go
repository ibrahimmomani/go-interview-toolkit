package collections

import (
	"fmt"
	"strings"
)

// Stack represents a Last-In-First-Out (LIFO) data structure with generic type support.
// Implemented using a slice for O(1) amortized operations.
type Stack[T any] struct {
	items []T
}

// NewStack creates and returns a new empty stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// NewStackWithCapacity creates a new stack with the specified initial capacity.
// This can help reduce memory allocations if you know the approximate size.
func NewStackWithCapacity[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, capacity),
	}
}

// FromSliceStack creates a new stack from a slice.
// The first element of the slice becomes the bottom of the stack.
func FromSliceStack[T any](slice []T) *Stack[T] {
	items := make([]T, len(slice))
	copy(items, slice)
	return &Stack[T]{
		items: items,
	}
}

// Push adds an element to the top of the stack.
// Time complexity: O(1) amortized
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the top element from the stack.
// Returns an error if the stack is empty.
// Time complexity: O(1)
func (s *Stack[T]) Pop() (T, error) {
	var zero T

	if len(s.items) == 0 {
		return zero, fmt.Errorf("stack is empty")
	}

	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]

	return value, nil
}

// Peek returns the top element without removing it.
// Returns an error if the stack is empty.
// Time complexity: O(1)
func (s *Stack[T]) Peek() (T, error) {
	var zero T

	if len(s.items) == 0 {
		return zero, fmt.Errorf("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// Size returns the number of elements in the stack.
// Time complexity: O(1)
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// IsEmpty returns true if the stack is empty.
// Time complexity: O(1)
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Clear removes all elements from the stack.
// Time complexity: O(1)
func (s *Stack[T]) Clear() {
	s.items = s.items[:0] // Keep the underlying array but set length to 0
}

// ToSlice returns a copy of the stack as a slice.
// The first element is the bottom of the stack, last element is the top.
// Time complexity: O(n)
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// Contains checks if the stack contains the specified value.
// Time complexity: O(n)
func (s *Stack[T]) Contains(value T) bool {
	for _, item := range s.items {
		if isEqual(item, value) {
			return true
		}
	}
	return false
}

// String returns a string representation of the stack.
// Shows elements from bottom to top.
func (s *Stack[T]) String() string {
	if len(s.items) == 0 {
		return "Stack[]"
	}

	var sb strings.Builder
	sb.WriteString("Stack[")

	for i, item := range s.items {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", item))
	}

	sb.WriteString("] (top)")
	return sb.String()
}

// Clone creates a deep copy of the stack.
// Time complexity: O(n)
func (s *Stack[T]) Clone() *Stack[T] {
	clone := NewStackWithCapacity[T](len(s.items))
	clone.items = make([]T, len(s.items))
	copy(clone.items, s.items)
	return clone
}

// Capacity returns the current capacity of the underlying slice.
// This can be useful for memory optimization analysis.
func (s *Stack[T]) Capacity() int {
	return cap(s.items)
}

// MultiPush pushes multiple elements onto the stack.
// Elements are pushed in order, so the last element will be at the top.
// Time complexity: O(n) where n is the number of elements
func (s *Stack[T]) MultiPush(values ...T) {
	s.items = append(s.items, values...)
}

// MultiPop pops n elements from the stack and returns them in reverse order.
// The first element in the returned slice is the top of the stack.
// Returns an error if there aren't enough elements.
// Time complexity: O(n)
func (s *Stack[T]) MultiPop(n int) ([]T, error) {
	if n < 0 {
		return nil, fmt.Errorf("cannot pop negative number of elements: %d", n)
	}

	if n > len(s.items) {
		return nil, fmt.Errorf("cannot pop %d elements from stack of size %d", n, len(s.items))
	}

	if n == 0 {
		return []T{}, nil
	}

	// Get the elements in reverse order (top to bottom)
	result := make([]T, n)
	start := len(s.items) - n

	for i := 0; i < n; i++ {
		result[i] = s.items[start+n-1-i]
	}

	// Remove the elements from the stack
	s.items = s.items[:start]

	return result, nil
}

// PeekN returns the top n elements without removing them.
// The first element in the returned slice is the top of the stack.
// Returns an error if there aren't enough elements.
// Time complexity: O(n)
func (s *Stack[T]) PeekN(n int) ([]T, error) {
	if n < 0 {
		return nil, fmt.Errorf("cannot peek negative number of elements: %d", n)
	}

	if n > len(s.items) {
		return nil, fmt.Errorf("cannot peek %d elements from stack of size %d", n, len(s.items))
	}

	if n == 0 {
		return []T{}, nil
	}

	result := make([]T, n)
	start := len(s.items) - n

	for i := 0; i < n; i++ {
		result[i] = s.items[start+n-1-i]
	}

	return result, nil
}

// Reverse reverses the order of elements in the stack.
// The bottom becomes the top and vice versa.
// Time complexity: O(n)
func (s *Stack[T]) Reverse() {
	for i, j := 0, len(s.items)-1; i < j; i, j = i+1, j-1 {
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}
