package collections

import (
	"reflect"
	"testing"
)

func TestNewDeque(t *testing.T) {
	dq := NewDeque[int]()

	if dq.Size() != 0 {
		t.Errorf("expected size 0, got %d", dq.Size())
	}

	if !dq.IsEmpty() {
		t.Error("expected empty deque")
	}

	if dq.Capacity() < 4 {
		t.Errorf("expected minimum capacity 4, got %d", dq.Capacity())
	}
}

func TestNewDequeWithCapacity(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 10},
		{0, 4},  // Should default to 4
		{-1, 4}, // Should default to 4
	}

	for _, tt := range tests {
		dq := NewDequeWithCapacity[int](tt.input)
		if dq.Capacity() != tt.expected {
			t.Errorf("input %d: expected capacity %d, got %d", tt.input, tt.expected, dq.Capacity())
		}
	}
}

func TestFromSliceDeque(t *testing.T) {
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
			dq := FromSliceDeque(tt.input)
			result := dq.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPushFront(t *testing.T) {
	dq := NewDeque[int]()

	// Test pushing to empty deque
	dq.PushFront(1)
	if dq.Size() != 1 {
		t.Errorf("expected size 1, got %d", dq.Size())
	}

	front, err := dq.Front()
	if err != nil || front != 1 {
		t.Errorf("expected front=1, got %d, error=%v", front, err)
	}

	// Test pushing multiple elements to front
	dq.PushFront(2)
	dq.PushFront(3)

	expected := []int{3, 2, 1}
	result := dq.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPushBack(t *testing.T) {
	dq := NewDeque[int]()

	// Test pushing to empty deque
	dq.PushBack(1)
	if dq.Size() != 1 {
		t.Errorf("expected size 1, got %d", dq.Size())
	}

	back, err := dq.Back()
	if err != nil || back != 1 {
		t.Errorf("expected back=1, got %d, error=%v", back, err)
	}

	// Test multiple elements
	dq.PushBack(2)
	dq.PushFront(0)

	front, _ := dq.Front()
	back, _ = dq.Back()

	if front != 0 {
		t.Errorf("expected front=0, got %d", front)
	}

	if back != 2 {
		t.Errorf("expected back=2, got %d", back)
	}
}

func TestDequeGetSet(t *testing.T) {
	dq := FromSliceDeque([]int{10, 20, 30, 40})

	// Test Get
	tests := []struct {
		name      string
		index     int
		expected  int
		expectErr bool
	}{
		{"get first", 0, 10, false},
		{"get middle", 2, 30, false},
		{"get last", 3, 40, false},
		{"invalid negative", -1, 0, true},
		{"invalid large", 4, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := dq.Get(tt.index)

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

			if value != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, value)
			}
		})
	}

	// Test Set
	err := dq.Set(1, 99)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	value, _ := dq.Get(1)
	if value != 99 {
		t.Errorf("expected 99 after set, got %d", value)
	}

	// Test invalid set
	err = dq.Set(10, 100)
	if err == nil {
		t.Error("expected error for invalid index")
	}
}

func TestDequeContains(t *testing.T) {
	dq := FromSliceDeque([]int{1, 2, 3})

	if !dq.Contains(2) {
		t.Error("expected to contain 2")
	}

	if dq.Contains(4) {
		t.Error("expected not to contain 4")
	}
}

func TestDequeClear(t *testing.T) {
	dq := FromSliceDeque([]int{1, 2, 3})
	dq.Clear()

	if dq.Size() != 0 {
		t.Errorf("expected size 0, got %d", dq.Size())
	}

	if !dq.IsEmpty() {
		t.Error("expected empty deque")
	}
}

func TestDequeClone(t *testing.T) {
	original := FromSliceDeque([]int{1, 2, 3})
	clone := original.Clone()

	// Check that clone has same contents
	if !reflect.DeepEqual(original.ToSlice(), clone.ToSlice()) {
		t.Error("clone should have same contents as original")
	}

	// Check that they are independent
	clone.PushBack(4)

	if original.Size() == clone.Size() {
		t.Error("clone should be independent of original")
	}
}

func TestDequeReverse(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{"empty deque", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"multiple elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"odd elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dq := FromSliceDeque(tt.initial)
			dq.Reverse()
			result := dq.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	dq := FromSliceDeque([]int{1, 2, 3, 4, 5})

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"rotate right 1", 1, []int{5, 1, 2, 3, 4}},
		{"rotate right 2", 2, []int{4, 5, 1, 2, 3}},
		{"rotate left 1", -1, []int{2, 3, 4, 5, 1}},
		{"rotate left 2", -2, []int{3, 4, 5, 1, 2}},
		{"rotate full cycle", 5, []int{1, 2, 3, 4, 5}},
		{"rotate 0", 0, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset deque
			dq.Clear()
			for _, v := range []int{1, 2, 3, 4, 5} {
				dq.PushBack(v)
			}

			dq.Rotate(tt.n)
			result := dq.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAliases(t *testing.T) {
	dq := NewDeque[int]()

	// Test queue aliases
	dq.Enqueue(1)
	dq.Enqueue(2)

	val, err := dq.Dequeue()
	if err != nil || val != 1 {
		t.Errorf("expected dequeue=1, got %d, error=%v", val, err)
	}

	// Test stack aliases
	dq.Push(3)
	val, err = dq.Pop()
	if err != nil || val != 3 {
		t.Errorf("expected pop=3, got %d, error=%v", val, err)
	}

	// Test peek aliases
	dq.PushBack(4)
	dq.PushFront(5)

	front, _ := dq.PeekFront()
	back, _ := dq.PeekBack()

	if front != 5 || back != 4 {
		t.Errorf("expected front=5, back=4, got front=%d, back=%d", front, back)
	}
}

func TestCircularBufferBehavior(t *testing.T) {
	dq := NewDequeWithCapacity[int](4)

	// Fill and test wraparound
	dq.PushBack(1)
	dq.PushBack(2)
	dq.PushFront(0)
	dq.PushFront(-1)

	expected := []int{-1, 0, 1, 2}
	result := dq.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Should trigger resize on next operation
	originalCap := dq.Capacity()
	dq.PushBack(3)

	if dq.Capacity() <= originalCap {
		t.Error("expected capacity to increase")
	}

	expected = []int{-1, 0, 1, 2, 3}
	result = dq.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestDequeString(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected string
	}{
		{"empty deque", []int{}, "Deque[]"},
		{"single element", []int{1}, "Deque[1] (front -> back)"},
		{"multiple elements", []int{1, 2, 3}, "Deque[1, 2, 3] (front -> back)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dq := FromSliceDeque(tt.initial)
			result := dq.String()

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestDequeAsStackAndQueue(t *testing.T) {
	// Test as stack (LIFO)
	stack := NewDeque[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val1, _ := stack.Pop()
	val2, _ := stack.Pop()
	val3, _ := stack.Pop()

	expected := []int{3, 2, 1}
	actual := []int{val1, val2, val3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("stack behavior: expected %v, got %v", expected, actual)
	}

	// Test as queue (FIFO)
	queue := NewDeque[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val1, _ = queue.Dequeue()
	val2, _ = queue.Dequeue()
	val3, _ = queue.Dequeue()

	expected = []int{1, 2, 3}
	actual = []int{val1, val2, val3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("queue behavior: expected %v, got %v", expected, actual)
	}
}

// Benchmark tests
func BenchmarkPushFront(b *testing.B) {
	dq := NewDeque[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dq.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	dq := NewDeque[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dq.PushBack(i)
	}
}

func BenchmarkPopFront(b *testing.B) {
	dq := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		dq.PushBack(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dq.PopFront()
	}
}

func BenchmarkPopBack(b *testing.B) {
	dq := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		dq.PushBack(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dq.PopBack()
	}
}

func BenchmarkMixedOperations(b *testing.B) {
	dq := NewDeque[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dq.PushFront(i)
		dq.PushBack(i)
		if i%2 == 0 {
			dq.PopFront()
		} else {
			dq.PopBack()
		}
	}
}
func TestMixedPushOperations(t *testing.T) {
	dq := NewDeque[int]()

	dq.PushBack(2)
	dq.PushFront(1)
	dq.PushBack(3)
	dq.PushFront(0)

	expected := []int{0, 1, 2, 3}
	result := dq.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestPopFront(t *testing.T) {
	dq := FromSliceDeque([]int{1, 2, 3})

	// Test normal pop
	value, err := dq.PopFront()
	if err != nil || value != 1 {
		t.Errorf("expected value=1, got %d, error=%v", value, err)
	}

	if dq.Size() != 2 {
		t.Errorf("expected size 2, got %d", dq.Size())
	}

	// Front should now be 2
	front, _ := dq.Front()
	if front != 2 {
		t.Errorf("expected front=2, got %d", front)
	}

	// Test pop until empty
	dq.PopFront() // 2
	dq.PopFront() // 3

	if !dq.IsEmpty() {
		t.Error("expected empty deque")
	}

	// Test pop from empty deque
	_, err = dq.PopFront()
	if err == nil {
		t.Error("expected error when popping from empty deque")
	}
}

func TestPopBack(t *testing.T) {
	dq := FromSliceDeque([]int{1, 2, 3})

	// Test normal pop
	value, err := dq.PopBack()
	if err != nil || value != 3 {
		t.Errorf("expected value=3, got %d, error=%v", value, err)
	}

	if dq.Size() != 2 {
		t.Errorf("expected size 2, got %d", dq.Size())
	}

	// Back should now be 2
	back, _ := dq.Back()
	if back != 2 {
		t.Errorf("expected back=2, got %d", back)
	}

	// Test pop until empty
	dq.PopBack() // 2
	dq.PopBack() // 1

	if !dq.IsEmpty() {
		t.Error("expected empty deque")
	}

	// Test pop from empty deque
	_, err = dq.PopBack()
	if err == nil {
		t.Error("expected error when popping from empty deque")
	}
}

func TestMixedPopOperations(t *testing.T) {
	dq := FromSliceDeque([]int{1, 2, 3, 4, 5})

	// Pop from both ends
	front, _ := dq.PopFront() // 1
	back, _ := dq.PopBack()   // 5

	if front != 1 || back != 5 {
		t.Errorf("expected front=1, back=5, got front=%d, back=%d", front, back)
	}

	expected := []int{2, 3, 4}
	result := dq.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFrontBack(t *testing.T) {
	// Test empty deque
	dq := NewDeque[int]()

	_, err := dq.Front()
	if err == nil {
		t.Error("expected error for empty deque front")
	}

	_, err = dq.Back()
	if err == nil {
		t.Error("expected error for empty deque back")
	}

	// Test single element
	dq.PushBack(1)
	front, err := dq.Front()
	if err != nil || front != 1 {
		t.Errorf("expected front=1, got %d, error=%v", front, err)
	}

	back, err := dq.Back()
	if err != nil || back != 1 {
		t.Errorf("expected back=1, got %d, error=%v", back, err)
	}

	// Test multiple elements
	dq.PushBack(2)
	dq.PushFront(0)

	front, _ = dq.Front()
	back, _ = dq.Back()

	if front != 0 {
		t.Errorf("expected front=0, got %d", front)
	}

	if back != 2 {
		t.Errorf("expected back=2, got %d", back)
	}
}
