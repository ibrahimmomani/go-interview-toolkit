package collections

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()

	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	if !s.IsEmpty() {
		t.Error("expected empty stack")
	}
}

func TestNewStackWithCapacity(t *testing.T) {
	s := NewStackWithCapacity[int](10)

	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	if s.Capacity() != 10 {
		t.Errorf("expected capacity 10, got %d", s.Capacity())
	}
}

func TestFromSliceStack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := FromSliceStack(tt.input)
			result := s.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPush(t *testing.T) {
	s := NewStack[int]()

	// Test pushing single element
	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("expected size 1, got %d", s.Size())
	}

	top, err := s.Peek()
	if err != nil || top != 1 {
		t.Errorf("expected top=1, got %d, error=%v", top, err)
	}

	// Test pushing multiple elements
	s.Push(2)
	s.Push(3)

	expected := []int{1, 2, 3}
	result := s.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Top should be the last pushed element
	top, _ = s.Peek()
	if top != 3 {
		t.Errorf("expected top=3, got %d", top)
	}
}

func TestPop(t *testing.T) {
	s := FromSliceStack([]int{1, 2, 3})

	// Test normal pop
	value, err := s.Pop()
	if err != nil || value != 3 {
		t.Errorf("expected value=3, got %d, error=%v", value, err)
	}

	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}

	// Test pop until empty
	s.Pop() // 2
	s.Pop() // 1

	if !s.IsEmpty() {
		t.Error("expected empty stack")
	}

	// Test pop from empty stack
	_, err = s.Pop()
	if err == nil {
		t.Error("expected error when popping from empty stack")
	}
}

func TestPeek(t *testing.T) {
	// Test peek on empty stack
	s := NewStack[int]()
	_, err := s.Peek()
	if err == nil {
		t.Error("expected error when peeking empty stack")
	}

	// Test peek with elements
	s.Push(1)
	s.Push(2)

	value, err := s.Peek()
	if err != nil || value != 2 {
		t.Errorf("expected value=2, got %d, error=%v", value, err)
	}

	// Size should remain unchanged
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}
}

func TestMultiPush(t *testing.T) {
	s := NewStack[int]()

	s.MultiPush(1, 2, 3, 4)

	expected := []int{1, 2, 3, 4}
	result := s.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Top should be the last element
	top, _ := s.Peek()
	if top != 4 {
		t.Errorf("expected top=4, got %d", top)
	}
}

func TestMultiPop(t *testing.T) {
	s := FromSliceStack([]int{1, 2, 3, 4, 5})

	tests := []struct {
		name      string
		n         int
		expected  []int
		expectErr bool
		remaining int
	}{
		{"pop 0 elements", 0, []int{}, false, 5},
		{"pop 2 elements", 2, []int{5, 4}, false, 3},
		{"pop remaining", 3, []int{3, 2, 1}, false, 0},
		{"pop from empty", 1, nil, true, 0},
		{"pop negative", -1, nil, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := s.MultiPop(tt.n)

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

			if s.Size() != tt.remaining {
				t.Errorf("expected remaining size %d, got %d", tt.remaining, s.Size())
			}
		})
	}
}

func TestPeekN(t *testing.T) {
	s := FromSliceStack([]int{1, 2, 3, 4, 5})

	tests := []struct {
		name      string
		n         int
		expected  []int
		expectErr bool
	}{
		{"peek 0 elements", 0, []int{}, false},
		{"peek 1 element", 1, []int{5}, false},
		{"peek 3 elements", 3, []int{5, 4, 3}, false},
		{"peek all elements", 5, []int{5, 4, 3, 2, 1}, false},
		{"peek too many", 6, nil, true},
		{"peek negative", -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := s.PeekN(tt.n)

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
			if s.Size() != 5 {
				t.Errorf("expected size 5, got %d", s.Size())
			}
		})
	}
}

func TestStackContains(t *testing.T) {
	s := FromSliceStack([]int{1, 2, 3})

	if !s.Contains(2) {
		t.Error("expected to contain 2")
	}

	if s.Contains(4) {
		t.Error("expected not to contain 4")
	}
}

func TestStackClear(t *testing.T) {
	s := FromSliceStack([]int{1, 2, 3})
	s.Clear()

	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	if !s.IsEmpty() {
		t.Error("expected empty stack")
	}
}

func TestStackClone(t *testing.T) {
	original := FromSliceStack([]int{1, 2, 3})
	clone := original.Clone()

	// Check that clone has same contents
	if !reflect.DeepEqual(original.ToSlice(), clone.ToSlice()) {
		t.Error("clone should have same contents as original")
	}

	// Check that they are independent
	clone.Push(4)

	if original.Size() == clone.Size() {
		t.Error("clone should be independent of original")
	}
}

func TestStackReverse(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{"empty stack", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"multiple elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := FromSliceStack(tt.initial)
			s.Reverse()
			result := s.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestStackString(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected string
	}{
		{"empty stack", []int{}, "Stack[]"},
		{"single element", []int{1}, "Stack[1] (top)"},
		{"multiple elements", []int{1, 2, 3}, "Stack[1, 2, 3] (top)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := FromSliceStack(tt.initial)
			result := s.String()

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestStackOperationsOrder(t *testing.T) {
	s := NewStack[int]()

	// Test LIFO behavior
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Should pop in reverse order
	val1, _ := s.Pop()
	val2, _ := s.Pop()
	val3, _ := s.Pop()

	expected := []int{3, 2, 1}
	actual := []int{val1, val2, val3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected LIFO order %v, got %v", expected, actual)
	}
}

// Benchmark tests
func BenchmarkPush(b *testing.B) {
	s := NewStack[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := NewStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	s := NewStack[int]()
	s.Push(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}

func BenchmarkStackWithCapacity(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := NewStackWithCapacity[int](1000)
		for j := 0; j < 1000; j++ {
			s.Push(j)
		}
	}
}
