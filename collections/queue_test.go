package collections

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()

	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}

	if !q.IsEmpty() {
		t.Error("expected empty queue")
	}

	if q.Capacity() < 4 {
		t.Errorf("expected minimum capacity 4, got %d", q.Capacity())
	}
}

func TestNewQueueWithCapacity(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 10},
		{0, 4},  // Should default to 4
		{-1, 4}, // Should default to 4
	}

	for _, tt := range tests {
		q := NewQueueWithCapacity[int](tt.input)
		if q.Capacity() != tt.expected {
			t.Errorf("input %d: expected capacity %d, got %d", tt.input, tt.expected, q.Capacity())
		}
	}
}

func TestFromSliceQueue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := FromSliceQueue(tt.input)
			result := q.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()

	// Test enqueuing single element
	q.Enqueue(1)
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	front, err := q.Front()
	if err != nil || front != 1 {
		t.Errorf("expected front=1, got %d, error=%v", front, err)
	}

	// Test enqueuing multiple elements
	q.Enqueue(2)
	q.Enqueue(3)

	expected := []int{1, 2, 3}
	result := q.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Front should remain the same, rear should be the last enqueued
	front, _ = q.Front()
	rear, _ := q.Rear()
	if front != 1 || rear != 3 {
		t.Errorf("expected front=1, rear=3, got front=%d, rear=%d", front, rear)
	}
}

func TestDequeue(t *testing.T) {
	q := FromSliceQueue([]int{1, 2, 3})

	// Test normal dequeue
	value, err := q.Dequeue()
	if err != nil || value != 1 {
		t.Errorf("expected value=1, got %d, error=%v", value, err)
	}

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	// Front should now be 2
	front, _ := q.Front()
	if front != 2 {
		t.Errorf("expected front=2, got %d", front)
	}

	// Test dequeue until empty
	q.Dequeue() // 2
	q.Dequeue() // 3

	if !q.IsEmpty() {
		t.Error("expected empty queue")
	}

	// Test dequeue from empty queue
	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue")
	}
}

func TestFrontRear(t *testing.T) {
	// Test empty queue
	q := NewQueue[int]()

	_, err := q.Front()
	if err == nil {
		t.Error("expected error for empty queue front")
	}

	_, err = q.Rear()
	if err == nil {
		t.Error("expected error for empty queue rear")
	}

	// Test single element
	q.Enqueue(1)
	front, err := q.Front()
	if err != nil || front != 1 {
		t.Errorf("expected front=1, got %d, error=%v", front, err)
	}

	rear, err := q.Rear()
	if err != nil || rear != 1 {
		t.Errorf("expected rear=1, got %d, error=%v", rear, err)
	}

	// Test multiple elements
	q.Enqueue(2)
	q.Enqueue(3)

	front, _ = q.Front()
	rear, _ = q.Rear()

	if front != 1 {
		t.Errorf("expected front=1, got %d", front)
	}

	if rear != 3 {
		t.Errorf("expected rear=3, got %d", rear)
	}
}

func TestCircularBuffer(t *testing.T) {
	// Test that the circular buffer works correctly
	q := NewQueueWithCapacity[int](4)

	// Fill to capacity
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	// Should trigger resize on next enqueue
	originalCap := q.Capacity()
	q.Enqueue(5)

	if q.Capacity() <= originalCap {
		t.Error("expected capacity to increase after exceeding limit")
	}

	// Test FIFO order
	expected := []int{1, 2, 3, 4, 5}
	result := q.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Dequeue some elements to test circular wrapping
	q.Dequeue() // 1
	q.Dequeue() // 2
	q.Enqueue(6)
	q.Enqueue(7)

	expected = []int{3, 4, 5, 6, 7}
	result = q.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMultiEnqueue(t *testing.T) {
	q := NewQueue[int]()

	q.MultiEnqueue(1, 2, 3, 4)

	expected := []int{1, 2, 3, 4}
	result := q.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Check front and rear
	front, _ := q.Front()
	rear, _ := q.Rear()
	if front != 1 || rear != 4 {
		t.Errorf("expected front=1, rear=4, got front=%d, rear=%d", front, rear)
	}
}

func TestMultiDequeue(t *testing.T) {
	q := FromSliceQueue([]int{1, 2, 3, 4, 5})

	tests := []struct {
		name      string
		n         int
		expected  []int
		expectErr bool
		remaining int
	}{
		{"dequeue 0 elements", 0, []int{}, false, 5},
		{"dequeue 2 elements", 2, []int{1, 2}, false, 3},
		{"dequeue remaining", 3, []int{3, 4, 5}, false, 0},
		{"dequeue from empty", 1, nil, true, 0},
		{"dequeue negative", -1, nil, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := q.MultiDequeue(tt.n)

			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}

			if q.Size() != tt.remaining {
				t.Errorf("expected remaining size %d, got %d", tt.remaining, q.Size())
			}
		})
	}
}

func TestQueuePeekN(t *testing.T) {
	q := FromSliceQueue([]int{11, 22, 33, 44, 55})

	tests := []struct {
		name      string
		n         int
		expected  []int
		expectErr bool
	}{
		{"peek 0 elements", 0, []int{}, false},
		{"peek 1 element", 1, []int{11}, false},
		{"peek 3 elements", 3, []int{11, 22, 33}, false},
		{"peek all elements", 5, []int{11, 22, 33, 44, 55}, false},
		{"peek too many", 6, nil, true},
		{"peek negative", -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := q.PeekN(tt.n)

			if tt.expectErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}

			// Size should remain unchanged
			if q.Size() != 5 {
				t.Errorf("expected size 5, got %d", q.Size())
			}
		})
	}
}

func TestQueueContains(t *testing.T) {
	q := FromSliceQueue([]int{13, 23, 33})

	if !q.Contains(23) {
		t.Error("expected to contain 23")
	}

	if q.Contains(4) {
		t.Error("expected not to contain 4")
	}
}

func TestQueueClear(t *testing.T) {
	q := FromSliceQueue([]int{13, 23, 33})
	q.Clear()

	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}

	if !q.IsEmpty() {
		t.Error("expected empty queue")
	}
}

func TestQueueClone(t *testing.T) {
	original := FromSliceQueue([]int{1, 2, 3})
	clone := original.Clone()

	// Check that clone has same contents
	if !reflect.DeepEqual(original.ToSlice(), clone.ToSlice()) {
		t.Error("clone should have same contents as original")
	}

	// Check that they are independent
	clone.Enqueue(4)

	if original.Size() == clone.Size() {
		t.Error("clone should be independent of original")
	}
}

func TestQueueReverse(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{"empty queue", []int{}, []int{}},
		{"single element", []int{1000}, []int{1000}},
		{"two elements", []int{100, 2}, []int{2, 100}},
		{"multiple elements", []int{11, 22, 33, 44}, []int{44, 33, 22, 11}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := FromSliceQueue(tt.initial)
			q.Reverse()
			result := q.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDrainTo(t *testing.T) {
	q := FromSliceQueue([]int{1, 2, 3, 4})

	result := q.DrainTo()
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	if !q.IsEmpty() {
		t.Error("expected empty queue after drain")
	}
}

func TestQueueString(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected string
	}{
		{"empty queue", []int{}, "Queue[]"},
		{"single element", []int{111}, "Queue[111] (front -> rear)"},
		{"multiple elements", []int{11, 22, 33}, "Queue[11, 22, 33] (front -> rear)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := FromSliceQueue(tt.initial)
			result := q.String()

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestQueueFIFOBehavior(t *testing.T) {
	q := NewQueue[int]()

	// Test FIFO behavior
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Should dequeue in same order
	val1, _ := q.Dequeue()
	val2, _ := q.Dequeue()
	val3, _ := q.Dequeue()

	expected := []int{1, 2, 3}
	actual := []int{val1, val2, val3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected FIFO order %v, got %v", expected, actual)
	}
}

func TestResizing(t *testing.T) {
	q := NewQueueWithCapacity[int](4)
	initialCap := q.Capacity()

	// Fill beyond initial capacity to trigger growth
	for i := 0; i < 8; i++ {
		q.Enqueue(i)
	}

	if q.Capacity() <= initialCap {
		t.Error("expected capacity to grow when queue is full")
	}

	// Dequeue most elements to trigger shrinking
	for i := 0; i < 6; i++ {
		q.Dequeue()
	}

	// Capacity should shrink (but this depends on implementation details)
	// We just verify the queue still works correctly
	q.Enqueue(100)
	q.Enqueue(200)

	expected := []int{6, 7, 100, 200}
	result := q.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Benchmark tests
func BenchmarkEnqueue(b *testing.B) {
	q := NewQueue[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkFront(b *testing.B) {
	q := NewQueue[int]()
	q.Enqueue(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Front()
	}
}

func BenchmarkQueueWithCapacity(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q := NewQueueWithCapacity[int](1000)
		for j := 0; j < 1000; j++ {
			q.Enqueue(j)
		}
		for j := 0; j < 1000; j++ {
			q.Dequeue()
		}
	}
}
