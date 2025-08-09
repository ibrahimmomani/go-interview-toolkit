package collections

import (
	"fmt"
	"strings"
)

const (
	// DefaultInitialCapacity is the default initial capacity for queues
	DequeInitialCapacity = 4
	// ShrinkFactor determines when to shrink the queue (when size == capacity/ShrinkFactor)
	DequeShrinkFactor = 4
	// GrowthFactor determines how much to grow the capacity
	DequeGrowthFactor = 2
)

// Deque represents a double-ended queue with generic type support.
// Elements can be added or removed from both ends efficiently.
// Implemented using a circular buffer with dynamic resizing.
type Deque[T any] struct {
	items []T
	front int // Index of the front element
	rear  int // Index where the next rear element will be inserted
	size  int // Current number of elements
}

// NewDeque creates and returns a new empty deque.
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		items: make([]T, DequeInitialCapacity), // Start with small capacity
		front: 0,
		rear:  0,
		size:  0,
	}
}

// NewDequeWithCapacity creates a new deque with the specified initial capacity.
func NewDequeWithCapacity[T any](capacity int) *Deque[T] {
	if capacity < 1 {
		capacity = DequeInitialCapacity
	}
	return &Deque[T]{
		items: make([]T, capacity),
		front: 0,
		rear:  0,
		size:  0,
	}
}

// FromSliceDeque creates a new deque from a slice.
// The first element of the slice becomes the front of the deque.
func FromSliceDeque[T any](slice []T) *Deque[T] {
	capacity := len(slice)
	if capacity < DequeInitialCapacity {
		capacity = DequeInitialCapacity
	}

	dq := &Deque[T]{
		items: make([]T, capacity),
		front: 0,
		rear:  len(slice),
		size:  len(slice),
	}

	copy(dq.items, slice)
	return dq
}

// PushFront adds an element to the front of the deque.
// Time complexity: O(1) amortized
func (dq *Deque[T]) PushFront(value T) {
	if dq.size == len(dq.items) {
		dq.resize()
	}

	dq.front = (dq.front - 1 + len(dq.items)) % len(dq.items)
	dq.items[dq.front] = value
	dq.size++
}

// PushBack adds an element to the back of the deque.
// Time complexity: O(1) amortized
func (dq *Deque[T]) PushBack(value T) {
	if dq.size == len(dq.items) {
		dq.resize()
	}

	dq.items[dq.rear] = value
	dq.rear = (dq.rear + 1) % len(dq.items)
	dq.size++
}

// PopFront removes and returns the front element from the deque.
// Returns an error if the deque is empty.
// Time complexity: O(1)
func (dq *Deque[T]) PopFront() (T, error) {
	var zero T

	if dq.size == 0 {
		return zero, fmt.Errorf("deque is empty")
	}

	value := dq.items[dq.front]
	var zeroVal T
	dq.items[dq.front] = zeroVal // Clear reference for GC
	dq.front = (dq.front + 1) % len(dq.items)
	dq.size--

	// Shrink if deque is 1/4 full and capacity > 4
	if dq.size > 0 && dq.size == len(dq.items)/DequeShrinkFactor && len(dq.items) > DequeShrinkFactor {
		dq.resize()
	}

	return value, nil
}

// PopBack removes and returns the back element from the deque.
// Returns an error if the deque is empty.
// Time complexity: O(1)
func (dq *Deque[T]) PopBack() (T, error) {
	var zero T

	if dq.size == 0 {
		return zero, fmt.Errorf("deque is empty")
	}

	dq.rear = (dq.rear - 1 + len(dq.items)) % len(dq.items)
	value := dq.items[dq.rear]
	var zeroVal T
	dq.items[dq.rear] = zeroVal // Clear reference for GC
	dq.size--

	// Shrink if deque is 1/4 full and capacity > 4
	if dq.size > 0 && dq.size == len(dq.items)/DequeShrinkFactor && len(dq.items) > DequeShrinkFactor {
		dq.resize()
	}

	return value, nil
}

// Front returns the front element without removing it.
// Returns an error if the deque is empty.
// Time complexity: O(1)
func (dq *Deque[T]) Front() (T, error) {
	var zero T

	if dq.size == 0 {
		return zero, fmt.Errorf("deque is empty")
	}

	return dq.items[dq.front], nil
}

// Back returns the back element without removing it.
// Returns an error if the deque is empty.
// Time complexity: O(1)
func (dq *Deque[T]) Back() (T, error) {
	var zero T

	if dq.size == 0 {
		return zero, fmt.Errorf("deque is empty")
	}

	backIndex := (dq.rear - 1 + len(dq.items)) % len(dq.items)
	return dq.items[backIndex], nil
}

// Size returns the number of elements in the deque.
// Time complexity: O(1)
func (dq *Deque[T]) Size() int {
	return dq.size
}

// IsEmpty returns true if the deque is empty.
// Time complexity: O(1)
func (dq *Deque[T]) IsEmpty() bool {
	return dq.size == 0
}

// Clear removes all elements from the deque.
// Time complexity: O(1)
func (dq *Deque[T]) Clear() {
	var zero T
	// Clear references for GC
	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % len(dq.items)
		dq.items[index] = zero
	}
	dq.front = 0
	dq.rear = 0
	dq.size = 0
}

// ToSlice returns a copy of the deque as a slice.
// The first element is the front of the deque.
// Time complexity: O(n)
func (dq *Deque[T]) ToSlice() []T {
	result := make([]T, dq.size)

	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % len(dq.items)
		result[i] = dq.items[index]
	}

	return result
}

// Contains checks if the deque contains the specified value.
// Time complexity: O(n)
func (dq *Deque[T]) Contains(value T) bool {
	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % len(dq.items)
		if isEqual(dq.items[index], value) {
			return true
		}
	}
	return false
}

// String returns a string representation of the deque.
// Shows elements from front to back.
func (dq *Deque[T]) String() string {
	if dq.size == 0 {
		return "Deque[]"
	}

	var sb strings.Builder
	sb.WriteString("Deque[")

	for i := 0; i < dq.size; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		index := (dq.front + i) % len(dq.items)
		sb.WriteString(fmt.Sprintf("%v", dq.items[index]))
	}

	sb.WriteString("] (front -> back)")
	return sb.String()
}

// Clone creates a deep copy of the deque.
// Time complexity: O(n)
func (dq *Deque[T]) Clone() *Deque[T] {
	clone := NewDequeWithCapacity[T](len(dq.items))

	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % len(dq.items)
		clone.PushBack(dq.items[index])
	}

	return clone
}

// Capacity returns the current capacity of the underlying slice.
func (dq *Deque[T]) Capacity() int {
	return len(dq.items)
}

// Get returns the element at the specified index (0 is front).
// Time complexity: O(1)
func (dq *Deque[T]) Get(index int) (T, error) {
	var zero T

	if index < 0 || index >= dq.size {
		return zero, fmt.Errorf("index %d out of bounds for deque of size %d", index, dq.size)
	}

	actualIndex := (dq.front + index) % len(dq.items)
	return dq.items[actualIndex], nil
}

// Set sets the element at the specified index (0 is front).
// Time complexity: O(1)
func (dq *Deque[T]) Set(index int, value T) error {
	if index < 0 || index >= dq.size {
		return fmt.Errorf("index %d out of bounds for deque of size %d", index, dq.size)
	}

	actualIndex := (dq.front + index) % len(dq.items)
	dq.items[actualIndex] = value
	return nil
}

// Reverse reverses the order of elements in the deque.
// Time complexity: O(n)
func (dq *Deque[T]) Reverse() {
	if dq.size <= 1 {
		return
	}

	// Swap elements from both ends moving inward
	for i := 0; i < dq.size/DequeGrowthFactor; i++ {
		frontIndex := (dq.front + i) % len(dq.items)
		backIndex := (dq.front + dq.size - 1 - i) % len(dq.items)
		dq.items[frontIndex], dq.items[backIndex] = dq.items[backIndex], dq.items[frontIndex]
	}
}

// resize doubles the capacity when full, halves when 1/4 full
func (dq *Deque[T]) resize() {
	var newCapacity int
	if dq.size == len(dq.items) {
		// Double when full
		newCapacity = len(dq.items) * DequeGrowthFactor
	} else {
		// Halve when 1/4 full
		newCapacity = len(dq.items) / DequeGrowthFactor
	}

	newItems := make([]T, newCapacity)

	// Copy elements in order
	for i := 0; i < dq.size; i++ {
		index := (dq.front + i) % len(dq.items)
		newItems[i] = dq.items[index]
	}

	dq.items = newItems
	dq.front = 0
	dq.rear = dq.size
}

// Rotate rotates the deque n positions to the right.
// Negative n rotates to the left.
// Time complexity: O(1)
func (dq *Deque[T]) Rotate(n int) {
	if dq.size <= 1 {
		return
	}

	// Normalize n to be within [-size, size]
	n %= dq.size
	if n < 0 {
		n += dq.size
	}

	// Adjust front pointer
	dq.front = (dq.front - n + len(dq.items)) % len(dq.items)
}

// PeekFront returns the front element (alias for Front).
func (dq *Deque[T]) PeekFront() (T, error) {
	return dq.Front()
}

// PeekBack returns the back element (alias for Back).
func (dq *Deque[T]) PeekBack() (T, error) {
	return dq.Back()
}

// Enqueue adds an element to the back (alias for PushBack for queue compatibility).
func (dq *Deque[T]) Enqueue(value T) {
	dq.PushBack(value)
}

// Dequeue removes and returns the front element (alias for PopFront for queue compatibility).
func (dq *Deque[T]) Dequeue() (T, error) {
	return dq.PopFront()
}

// Push adds an element to the back (alias for PushBack for stack compatibility).
func (dq *Deque[T]) Push(value T) {
	dq.PushBack(value)
}

// Pop removes and returns the back element (alias for PopBack for stack compatibility).
func (dq *Deque[T]) Pop() (T, error) {
	return dq.PopBack()
}
