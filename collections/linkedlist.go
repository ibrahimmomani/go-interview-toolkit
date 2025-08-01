// Package collections provides linear data structures for coding interviews.
package collections

import (
	"fmt"
	"strings"
)

// Node represents a single node in the linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a singly linked list with generic type support.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewLinkedList creates and returns a new empty linked list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// FromSlice creates a new linked list from a slice.
func FromSlice[T any](slice []T) *LinkedList[T] {
	ll := NewLinkedList[T]()
	for _, v := range slice {
		ll.Append(v)
	}
	return ll
}

// Append adds an element to the end of the list.
// Time complexity: O(1)
func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.Next = newNode
		ll.tail = newNode
	}

	ll.size++
}

// Prepend adds an element to the beginning of the list.
// Time complexity: O(1)
func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: ll.head}
	ll.head = newNode

	if ll.tail == nil {
		ll.tail = newNode
	}

	ll.size++
}

// Insert adds an element at the specified index.
// Time complexity: O(n)
func (ll *LinkedList[T]) Insert(index int, value T) error {
	if index < 0 || index > ll.size {
		return fmt.Errorf("index %d out of bounds for list of size %d", index, ll.size)
	}

	if index == 0 {
		ll.Prepend(value)
		return nil
	}

	if index == ll.size {
		ll.Append(value)
		return nil
	}

	newNode := &Node[T]{Value: value}
	current := ll.head

	// Navigate to position index-1
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	ll.size++

	return nil
}

// Delete removes the first occurrence of the specified value.
// Time complexity: O(n)
func (ll *LinkedList[T]) Delete(value T) bool {
	if ll.head == nil {
		return false
	}

	// Handle deletion of head node
	if isEqual(ll.head.Value, value) {
		ll.head = ll.head.Next
		if ll.head == nil {
			ll.tail = nil
		}
		ll.size--
		return true
	}

	current := ll.head
	for current.Next != nil {
		if isEqual(current.Next.Value, value) {
			nodeToDelete := current.Next
			current.Next = nodeToDelete.Next

			// Update tail if we deleted the last node
			if nodeToDelete == ll.tail {
				ll.tail = current
			}

			ll.size--
			return true
		}
		current = current.Next
	}

	return false
}

// DeleteAt removes the element at the specified index.
// Time complexity: O(n)
func (ll *LinkedList[T]) DeleteAt(index int) error {
	if index < 0 || index >= ll.size {
		return fmt.Errorf("index %d out of bounds for list of size %d", index, ll.size)
	}

	// Handle deletion of head node
	if index == 0 {
		ll.head = ll.head.Next
		if ll.head == nil {
			ll.tail = nil
		}
		ll.size--
		return nil
	}

	current := ll.head
	// Navigate to position index-1
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	nodeToDelete := current.Next
	current.Next = nodeToDelete.Next

	// Update tail if we deleted the last node
	if nodeToDelete == ll.tail {
		ll.tail = current
	}

	ll.size--
	return nil
}

// Get returns the element at the specified index.
// Time complexity: O(n)
func (ll *LinkedList[T]) Get(index int) (T, error) {
	var zero T

	if index < 0 || index >= ll.size {
		return zero, fmt.Errorf("index %d out of bounds for list of size %d", index, ll.size)
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, nil
}

// Find returns the index of the first occurrence of the specified value.
// Returns -1 if not found.
// Time complexity: O(n)
func (ll *LinkedList[T]) Find(value T) int {
	current := ll.head
	index := 0

	for current != nil {
		if isEqual(current.Value, value) {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

// Contains checks if the list contains the specified value.
// Time complexity: O(n)
func (ll *LinkedList[T]) Contains(value T) bool {
	return ll.Find(value) != -1
}

// Size returns the number of elements in the list.
// Time complexity: O(1)
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// IsEmpty returns true if the list is empty.
// Time complexity: O(1)
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

// Clear removes all elements from the list.
// Time complexity: O(1)
func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

// Head returns the first element without removing it.
// Time complexity: O(1)
func (ll *LinkedList[T]) Head() (T, error) {
	var zero T

	if ll.head == nil {
		return zero, fmt.Errorf("list is empty")
	}

	return ll.head.Value, nil
}

// Tail returns the last element without removing it.
// Time complexity: O(1)
func (ll *LinkedList[T]) Tail() (T, error) {
	var zero T

	if ll.tail == nil {
		return zero, fmt.Errorf("list is empty")
	}

	return ll.tail.Value, nil
}

// ToSlice converts the linked list to a slice.
// Time complexity: O(n)
func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, ll.size)
	current := ll.head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// Reverse reverses the linked list in place.
// Time complexity: O(n)
func (ll *LinkedList[T]) Reverse() {
	if ll.head == nil || ll.head.Next == nil {
		return
	}

	var prev *Node[T]
	current := ll.head
	ll.tail = ll.head // The current head will become the tail

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.head = prev
}

// String returns a string representation of the linked list.
func (ll *LinkedList[T]) String() string {
	if ll.size == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	current := ll.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.Value))
		if current.Next != nil {
			sb.WriteString(" -> ")
		}
		current = current.Next
	}

	sb.WriteString("]")
	return sb.String()
}

// GetNode returns the node at the specified index (useful for advanced operations).
// Time complexity: O(n)
func (ll *LinkedList[T]) GetNode(index int) (*Node[T], error) {
	if index < 0 || index >= ll.size {
		return nil, fmt.Errorf("index %d out of bounds for list of size %d", index, ll.size)
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current, nil
}

// isEqual compares two values for equality using fmt.Sprintf for comparison.
// This works for most types but can be overridden for custom comparison logic.
func isEqual[T any](a, b T) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
