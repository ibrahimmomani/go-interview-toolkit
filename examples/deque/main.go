package main

import (
	"fmt"

	"github.com/ibrahimmomani/go-interview-toolkit/collections"
)

func main() {
	fmt.Println("=== Deque (Double-ended Queue) Demo ===")

	// Create a new deque
	deque := collections.NewDeque[int]()
	fmt.Println("Created empty deque:", deque)

	// Push elements to both ends
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushFront(1)
	deque.PushFront(0)
	fmt.Println("After pushing to both ends:", deque)

	// Check front and back
	if front, err := deque.Front(); err == nil {
		fmt.Printf("Front element: %d\n", front)
	}
	if back, err := deque.Back(); err == nil {
		fmt.Printf("Back element: %d\n", back)
	}

	// Pop from both ends
	if frontVal, err := deque.PopFront(); err == nil {
		fmt.Printf("Popped front: %d, Deque now: %s\n", frontVal, deque)
	}
	if backVal, err := deque.PopBack(); err == nil {
		fmt.Printf("Popped back: %d, Deque now: %s\n", backVal, deque)
	}

	// Access by index
	deque.PushBack(4)
	deque.PushBack(5)
	fmt.Println("After adding more elements:", deque)

	if val, err := deque.Get(1); err == nil {
		fmt.Printf("Element at index 1: %d\n", val)
	}

	// Set element by index
	deque.Set(0, 99)
	fmt.Println("After setting index 0 to 99:", deque)

	// Rotate the deque
	fmt.Println("\nRotation examples:")
	deque.Clear()
	for i := 1; i <= 5; i++ {
		deque.PushBack(i)
	}
	fmt.Println("Original:", deque)

	clone1 := deque.Clone()
	clone1.Rotate(2)
	fmt.Println("Rotated right by 2:", clone1)

	clone2 := deque.Clone()
	clone2.Rotate(-1)
	fmt.Println("Rotated left by 1:", clone2)

	// Reverse the deque
	deque.Reverse()
	fmt.Println("Reversed:", deque)

	// Demonstrate aliases (can work as stack or queue)
	fmt.Println("\n=== Using Deque as Stack ===")
	stack := collections.NewDeque[string]()
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	fmt.Println("Stack after pushes:", stack)

	for !stack.IsEmpty() {
		if val, err := stack.Pop(); err == nil {
			fmt.Printf("Popped: %s\n", val)
		}
	}

	fmt.Println("\n=== Using Deque as Queue ===")
	queue := collections.NewDeque[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	fmt.Println("Queue after enqueues:", queue)

	for !queue.IsEmpty() {
		if val, err := queue.Dequeue(); err == nil {
			fmt.Printf("Dequeued: %s\n", val)
		}
	}

	// Demonstrate common interview use cases
	fmt.Println("\n=== Interview Use Cases ===")

	// 1. Sliding Window Maximum
	fmt.Println("1. Sliding Window Maximum:")
	slidingWindowMaximum([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)

	// 2. Palindrome checker
	fmt.Println("\n2. Palindrome Checker:")
	words := []string{"racecar", "hello", "level", "world", "madam"}
	for _, word := range words {
		isPalin := isPalindrome(word)
		fmt.Printf("'%s' is palindrome: %t\n", word, isPalin)
	}

	// 3. Maximum in all subarrays of size k
	fmt.Println("\n3. Maximum in All Subarrays:")
	maxInSubarrays([]int{12, 1, 78, 90, 57, 89, 56}, 3)

	// 4. First negative in every window
	fmt.Println("\n4. First Negative in Every Window:")
	firstNegativeInWindow([]int{12, -1, -7, 8, -15, 30, 16, 28}, 3)
}

// Sliding Window Maximum using deque - stores indices
func slidingWindowMaximum(arr []int, k int) {
	if len(arr) == 0 || k <= 0 {
		return
	}

	dq := collections.NewDeque[int]() // Store indices
	result := make([]int, 0)

	fmt.Printf("Array: %v, Window size: %d\n", arr, k)

	for i := 0; i < len(arr); i++ {
		// Remove indices that are out of current window
		for !dq.IsEmpty() {
			if front, _ := dq.Front(); front <= i-k {
				dq.PopFront()
			} else {
				break
			}
		}

		// Remove indices whose corresponding values are smaller than current element
		for !dq.IsEmpty() {
			if back, _ := dq.Back(); arr[back] <= arr[i] {
				dq.PopBack()
			} else {
				break
			}
		}

		// Add current element index
		dq.PushBack(i)

		// The front of deque contains the index of maximum element of current window
		if i >= k-1 {
			if front, _ := dq.Front(); front >= 0 {
				result = append(result, arr[front])
			}
		}
	}

	fmt.Printf("Maximum in each window: %v\n", result)
}

// Check if string is palindrome using deque
func isPalindrome(s string) bool {
	dq := collections.NewDeque[rune]()

	// Add all characters to deque
	for _, char := range s {
		dq.PushBack(char)
	}

	// Compare characters from both ends
	for dq.Size() > 1 {
		front, _ := dq.PopFront()
		back, _ := dq.PopBack()

		if front != back {
			return false
		}
	}

	return true
}

// Find maximum in all subarrays of size k
func maxInSubarrays(arr []int, k int) {
	if len(arr) < k {
		return
	}

	dq := collections.NewDeque[int]() // Store indices
	fmt.Printf("Array: %v, Subarray size: %d\n", arr, k)

	// Process first window
	for i := 0; i < k; i++ {
		// Remove smaller elements from rear
		for !dq.IsEmpty() {
			if back, _ := dq.Back(); arr[back] <= arr[i] {
				dq.PopBack()
			} else {
				break
			}
		}
		dq.PushBack(i)
	}

	// The front contains index of the largest element
	if front, _ := dq.Front(); front >= 0 {
		fmt.Printf("Maximum in subarray [0:%d]: %d\n", k, arr[front])
	}

	// Process remaining elements
	for i := k; i < len(arr); i++ {
		// Remove elements out of current window
		for !dq.IsEmpty() {
			if front, _ := dq.Front(); front <= i-k {
				dq.PopFront()
			} else {
				break
			}
		}

		// Remove smaller elements from rear
		for !dq.IsEmpty() {
			if back, _ := dq.Back(); arr[back] <= arr[i] {
				dq.PopBack()
			} else {
				break
			}
		}

		dq.PushBack(i)

		if front, _ := dq.Front(); front >= 0 {
			fmt.Printf("Maximum in subarray [%d:%d]: %d\n", i-k+1, i+1, arr[front])
		}
	}
}

// Find first negative number in every window of size k
func firstNegativeInWindow(arr []int, k int) {
	if len(arr) < k {
		return
	}

	dq := collections.NewDeque[int]() // Store indices of negative numbers
	fmt.Printf("Array: %v, Window size: %d\n", arr, k)

	// Process first window
	for i := 0; i < k; i++ {
		if arr[i] < 0 {
			dq.PushBack(i)
		}
	}

	// Print result for first window
	if !dq.IsEmpty() {
		if front, _ := dq.Front(); front >= 0 {
			fmt.Printf("Window [0:%d]: First negative = %d\n", k, arr[front])
		}
	} else {
		fmt.Printf("Window [0:%d]: No negative number\n", k)
	}

	// Process remaining windows
	for i := k; i < len(arr); i++ {
		// Remove indices that are out of current window
		for !dq.IsEmpty() {
			if front, _ := dq.Front(); front <= i-k {
				dq.PopFront()
			} else {
				break
			}
		}

		// Add current element if it's negative
		if arr[i] < 0 {
			dq.PushBack(i)
		}

		// Print first negative for current window
		if !dq.IsEmpty() {
			if front, _ := dq.Front(); front >= 0 {
				fmt.Printf("Window [%d:%d]: First negative = %d\n", i-k+1, i+1, arr[front])
			}
		} else {
			fmt.Printf("Window [%d:%d]: No negative number\n", i-k+1, i+1)
		}
	}
}
