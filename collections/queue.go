package collections

import (
	"fmt"
	"strings"
)

const (
	// DefaultInitialCapacity is the default initial capacity for queues
	DefaultInitialCapacity = 4
	// ShrinkFactor determines when to shrink the queue (when size == capacity/ShrinkFactor)
	ShrinkFactor = 4
	// GrowthFactor determines how much to grow the capacity
	GrowthFactor = 2
)

// Queue represents a First-In-First-Out (FIFO) data structure with generic type support.
// Implemented using a circular buffer with dynamic resizing for optimal performance.
type Queue[T any] struct {
	items []T
	front int // Index of the front element
	rear  int // Index where the next element will be inserted
	size  int // Current number of elements
}

// NewQueue creates and returns a new empty queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, DefaultInitialCapacity), // Start with small capacity
		front: 0,
		rear:  0,
		size:  0,
	}
}

// NewQueueWithCapacity creates a new queue with the specified initial capacity.
func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	if capacity < 1 {
		capacity = 4
	}
	return &Queue[T]{
		items: make([]T, capacity),
		front: 0,
		rear:  0,
		size:  0,
	}
}

// FromSliceQueue creates a new queue from a slice.
// The first element of the slice becomes the front of the queue.
func FromSliceQueue[T any](slice []T) *Queue[T] {
	capacity := len(slice)
	if capacity < DefaultInitialCapacity {
		capacity = DefaultInitialCapacity
	}

	q := &Queue[T]{
		items: make([]T, capacity),
		front: 0,
		rear:  len(slice),
		size:  len(slice),
	}

	copy(q.items, slice)
	return q
}

// Enqueue adds an element to the rear of the queue.
// Time complexity: O(1) amortized
func (q *Queue[T]) Enqueue(value T) {
	if q.size == len(q.items) {
		q.resize()
	}

	q.items[q.rear] = value
	q.rear = (q.rear + 1) % len(q.items)
	q.size++
}

// Dequeue removes and returns the front element from the queue.
// Returns an error if the queue is empty.
// Time complexity: O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T

	if q.size == 0 {
		return zero, fmt.Errorf("queue is empty")
	}

	value := q.items[q.front]
	var zeroVal T
	q.items[q.front] = zeroVal // Clear the reference for GC
	q.front = (q.front + 1) % len(q.items)
	q.size--

	// Shrink if queue is 1/4 full and capacity > 4
	if q.size > 0 && q.size == len(q.items)/ShrinkFactor && len(q.items) > ShrinkFactor {
		q.resize()
	}

	return value, nil
}

// Front returns the front element without removing it.
// Returns an error if the queue is empty.
// Time complexity: O(1)
func (q *Queue[T]) Front() (T, error) {
	var zero T

	if q.size == 0 {
		return zero, fmt.Errorf("queue is empty")
	}

	return q.items[q.front], nil
}

// Rear returns the rear element without removing it.
// Returns an error if the queue is empty.
// Time complexity: O(1)
func (q *Queue[T]) Rear() (T, error) {
	var zero T

	if q.size == 0 {
		return zero, fmt.Errorf("queue is empty")
	}

	rearIndex := (q.rear - 1 + len(q.items)) % len(q.items)
	return q.items[rearIndex], nil
}

// Size returns the number of elements in the queue.
// Time complexity: O(1)
func (q *Queue[T]) Size() int {
	return q.size
}

// IsEmpty returns true if the queue is empty.
// Time complexity: O(1)
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Clear removes all elements from the queue.
// Time complexity: O(1)
func (q *Queue[T]) Clear() {
	var zero T
	// Clear references for GC
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % len(q.items)
		q.items[index] = zero
	}
	q.front = 0
	q.rear = 0
	q.size = 0
}

// ToSlice returns a copy of the queue as a slice.
// The first element is the front of the queue.
// Time complexity: O(n)
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, q.size)

	for i := 0; i < q.size; i++ {
		index := (q.front + i) % len(q.items)
		result[i] = q.items[index]
	}

	return result
}

// Contains checks if the queue contains the specified value.
// Time complexity: O(n)
func (q *Queue[T]) Contains(value T) bool {
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % len(q.items)
		if isEqual(q.items[index], value) {
			return true
		}
	}
	return false
}

// String returns a string representation of the queue.
// Shows elements from front to rear.
func (q *Queue[T]) String() string {
	if q.size == 0 {
		return "Queue[]"
	}

	var sb strings.Builder
	sb.WriteString("Queue[")

	for i := 0; i < q.size; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		index := (q.front + i) % len(q.items)
		sb.WriteString(fmt.Sprintf("%v", q.items[index]))
	}

	sb.WriteString("] (front -> rear)")
	return sb.String()
}

// Clone creates a deep copy of the queue.
// Time complexity: O(n)
func (q *Queue[T]) Clone() *Queue[T] {
	clone := NewQueueWithCapacity[T](len(q.items))

	for i := 0; i < q.size; i++ {
		index := (q.front + i) % len(q.items)
		clone.Enqueue(q.items[index])
	}

	return clone
}

// Capacity returns the current capacity of the underlying slice.
func (q *Queue[T]) Capacity() int {
	return len(q.items)
}

// MultiEnqueue adds multiple elements to the rear of the queue.
// Time complexity: O(n) where n is the number of elements
func (q *Queue[T]) MultiEnqueue(values ...T) {
	for _, value := range values {
		q.Enqueue(value)
	}
}

// MultiDequeue removes n elements from the front of the queue.
// Returns the elements in the order they were dequeued.
// Returns an error if there aren't enough elements.
// Time complexity: O(n)
func (q *Queue[T]) MultiDequeue(n int) ([]T, error) {
	if n < 0 {
		return nil, fmt.Errorf("cannot dequeue negative number of elements: %d", n)
	}

	if n > q.size {
		return nil, fmt.Errorf("cannot dequeue %d elements from queue of size %d", n, q.size)
	}

	if n == 0 {
		return []T{}, nil
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		value, _ := q.Dequeue()
		result[i] = value
	}

	return result, nil
}

// PeekN returns the front n elements without removing them.
// Returns an error if there aren't enough elements.
// Time complexity: O(n)
func (q *Queue[T]) PeekN(n int) ([]T, error) {
	if n < 0 {
		return nil, fmt.Errorf("cannot peek negative number of elements: %d", n)
	}

	if n > q.size {
		return nil, fmt.Errorf("cannot peek %d elements from queue of size %d", n, q.size)
	}

	if n == 0 {
		return []T{}, nil
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		index := (q.front + i) % len(q.items)
		result[i] = q.items[index]
	}

	return result, nil
}

// Reverse reverses the order of elements in the queue.
// The front becomes the rear and vice versa.
// Time complexity: O(n)
func (q *Queue[T]) Reverse() {
	if q.size <= 1 {
		return
	}

	// Convert to slice, reverse it, then rebuild queue
	slice := q.ToSlice()

	// Reverse the slice
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	// Clear and rebuild
	q.Clear()
	for _, value := range slice {
		q.Enqueue(value)
	}
}

// resize doubles the capacity when full, halves when 1/4 full
func (q *Queue[T]) resize() {
	var newCapacity int
	if q.size == len(q.items) {
		// Double when full
		newCapacity = len(q.items) * GrowthFactor
	} else {
		// Halve when 1/4 full
		newCapacity = len(q.items) / GrowthFactor
	}

	newItems := make([]T, newCapacity)

	// Copy elements in order
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % len(q.items)
		newItems[i] = q.items[index]
	}

	q.items = newItems
	q.front = 0
	q.rear = q.size
}

// DrainTo removes all elements from the queue and returns them as a slice.
// The queue becomes empty after this operation.
// Time complexity: O(n)
func (q *Queue[T]) DrainTo() []T {
	result := q.ToSlice()
	q.Clear()
	return result
}

// Peek returns the front element (alias for Front for consistency with other collections).
func (q *Queue[T]) Peek() (T, error) {
	return q.Front()
}
