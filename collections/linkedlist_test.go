package collections

import (
	"reflect"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()

	if ll.Size() != 0 {
		t.Errorf("expected size 0, got %d", ll.Size())
	}

	if !ll.IsEmpty() {
		t.Error("expected empty list")
	}

	if ll.head != nil {
		t.Error("expected nil head")
	}

	if ll.tail != nil {
		t.Error("expected nil tail")
	}
}

func TestFromSlice(t *testing.T) {
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
			ll := FromSlice(tt.input)
			result := ll.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	ll := NewLinkedList[int]()

	// Test appending to empty list
	ll.Append(1)
	if ll.Size() != 1 {
		t.Errorf("expected size 1, got %d", ll.Size())
	}

	head, _ := ll.Head()
	tail, _ := ll.Tail()
	if head != 1 || tail != 1 {
		t.Errorf("expected head and tail to be 1, got head=%d, tail=%d", head, tail)
	}

	// Test appending multiple elements
	ll.Append(2)
	ll.Append(3)

	expected := []int{1, 2, 3}
	result := ll.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	tail, _ = ll.Tail()
	if tail != 3 {
		t.Errorf("expected tail to be 3, got %d", tail)
	}
}

func TestPrepend(t *testing.T) {
	ll := NewLinkedList[int]()

	// Test prepending to empty list
	ll.Prepend(1)
	if ll.Size() != 1 {
		t.Errorf("expected size 1, got %d", ll.Size())
	}

	// Test prepending multiple elements
	ll.Prepend(2)
	ll.Prepend(3)

	expected := []int{3, 2, 1}
	result := ll.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestInsert(t *testing.T) {
	ll := FromSlice([]int{1, 3, 5})

	tests := []struct {
		name      string
		index     int
		value     int
		expected  []int
		expectErr bool
	}{
		{"insert at beginning", 0, 0, []int{0, 1, 3, 5}, false},
		{"insert in middle", 2, 4, []int{0, 1, 4, 3, 5}, false},
		{"insert at end", 5, 6, []int{0, 1, 4, 3, 5, 6}, false},
		{"invalid negative index", -1, 10, nil, true},
		{"invalid large index", 10, 10, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ll.Insert(tt.index, tt.value)

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

			result := ll.ToSlice()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		delete   int
		expected []int
		found    bool
	}{
		{"delete from empty", []int{}, 1, []int{}, false},
		{"delete head", []int{1, 2, 3}, 1, []int{2, 3}, true},
		{"delete middle", []int{1, 2, 3}, 2, []int{1, 3}, true},
		{"delete tail", []int{1, 2, 3}, 3, []int{1, 2}, true},
		{"delete non-existent", []int{1, 2, 3}, 4, []int{1, 2, 3}, false},
		{"delete single element", []int{1}, 1, []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := FromSlice(tt.initial)
			found := ll.Delete(tt.delete)

			if found != tt.found {
				t.Errorf("expected found=%v, got %v", tt.found, found)
			}

			result := ll.ToSlice()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		index     int
		expected  []int
		expectErr bool
	}{
		{"delete at 0", []int{1, 2, 3}, 0, []int{2, 3}, false},
		{"delete in middle", []int{1, 2, 3}, 1, []int{1, 3}, false},
		{"delete at end", []int{1, 2, 3}, 2, []int{1, 2}, false},
		{"invalid negative index", []int{1, 2, 3}, -1, nil, true},
		{"invalid large index", []int{1, 2, 3}, 3, nil, true},
		{"delete from single element", []int{1}, 0, []int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := FromSlice(tt.initial)
			err := ll.DeleteAt(tt.index)

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

			result := ll.ToSlice()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestGet(t *testing.T) {
	ll := FromSlice([]int{10, 20, 30})

	tests := []struct {
		name      string
		index     int
		expected  int
		expectErr bool
	}{
		{"get first", 0, 10, false},
		{"get middle", 1, 20, false},
		{"get last", 2, 30, false},
		{"invalid negative", -1, 0, true},
		{"invalid large", 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := ll.Get(tt.index)

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
}

func TestFind(t *testing.T) {
	ll := FromSlice([]int{10, 20, 30, 20})

	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"find first occurrence", 20, 1},
		{"find at beginning", 10, 0},
		{"find at end", 30, 2},
		{"not found", 40, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := ll.Find(tt.value)
			if index != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, index)
			}
		})
	}
}

func TestContains(t *testing.T) {
	ll := FromSlice([]int{1, 2, 3})

	if !ll.Contains(2) {
		t.Error("expected to contain 2")
	}

	if ll.Contains(4) {
		t.Error("expected not to contain 4")
	}
}

func TestHeadTail(t *testing.T) {
	// Test empty list
	ll := NewLinkedList[int]()

	_, err := ll.Head()
	if err == nil {
		t.Error("expected error for empty list head")
	}

	_, err = ll.Tail()
	if err == nil {
		t.Error("expected error for empty list tail")
	}

	// Test single element
	ll.Append(1)
	head, err := ll.Head()
	if err != nil || head != 1 {
		t.Errorf("expected head=1, got %d, error=%v", head, err)
	}

	tail, err := ll.Tail()
	if err != nil || tail != 1 {
		t.Errorf("expected tail=1, got %d, error=%v", tail, err)
	}

	// Test multiple elements
	ll.Append(2)
	ll.Append(3)

	head, _ = ll.Head()
	tail, _ = ll.Tail()

	if head != 1 {
		t.Errorf("expected head=1, got %d", head)
	}

	if tail != 3 {
		t.Errorf("expected tail=3, got %d", tail)
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{"empty list", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"multiple elements", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := FromSlice(tt.initial)
			ll.Reverse()
			result := ll.ToSlice()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestClear(t *testing.T) {
	ll := FromSlice([]int{1, 2, 3})
	ll.Clear()

	if ll.Size() != 0 {
		t.Errorf("expected size 0, got %d", ll.Size())
	}

	if !ll.IsEmpty() {
		t.Error("expected empty list")
	}

	if ll.head != nil || ll.tail != nil {
		t.Error("expected nil head and tail")
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected string
	}{
		{"empty list", []int{}, "[]"},
		{"single element", []int{1}, "[1]"},
		{"multiple elements", []int{1, 2, 3}, "[1 -> 2 -> 3]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := FromSlice(tt.initial)
			result := ll.String()

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkAppend(b *testing.B) {
	ll := NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Append(i)
	}
}

func BenchmarkPrepend(b *testing.B) {
	ll := NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Prepend(i)
	}
}

func BenchmarkGet(b *testing.B) {
	ll := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Get(i % 1000)
	}
}

func BenchmarkFind(b *testing.B) {
	ll := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Find(i % 1000)
	}
}
