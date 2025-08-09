package main

import (
	"fmt"

	"github.com/ibrahimmomani/go-interview-toolkit/collections"
)

func main() {
	fmt.Println("=== Queue Demo ===")

	// Create a new queue
	queue := collections.NewQueue[int]()
	fmt.Println("Created empty queue:", queue)

	// Enqueue elements (FIFO - First In, First Out)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("After enqueuing 1, 2, 3:", queue)

	// Check front and rear
	if front, err := queue.Front(); err == nil {
		fmt.Printf("Front element: %d\n", front)
	}
	if rear, err := queue.Rear(); err == nil {
		fmt.Printf("Rear element: %d\n", rear)
	}

	// Dequeue elements
	if val, err := queue.Dequeue(); err == nil {
		fmt.Printf("Dequeued: %d, Queue now: %s\n", val, queue)
	}

	// Multi-enqueue
	queue.MultiEnqueue(10, 20, 30)
	fmt.Println("After multi-enqueue 10, 20, 30:", queue)

	// Multi-dequeue
	if values, err := queue.MultiDequeue(2); err == nil {
		fmt.Printf("Multi-dequeued: %v, Queue now: %s\n", values, queue)
	}

	// Peek multiple elements
	if values, err := queue.PeekN(3); err == nil {
		fmt.Printf("Front 3 elements (peek): %v\n", values)
	}

	// Convert to slice
	slice := queue.ToSlice()
	fmt.Printf("As slice: %v\n", slice)

	// Clone the queue
	clone := queue.Clone()
	fmt.Println("Cloned queue:", clone)

	// Reverse the original
	queue.Reverse()
	fmt.Println("After reversing original:", queue)
	fmt.Println("Clone remains unchanged:", clone)

	// Drain all elements
	drained := clone.DrainTo()
	fmt.Printf("Drained from clone: %v, Clone now: %s\n", drained, clone)

	// Create from slice
	newQueue := collections.FromSliceQueue([]string{"first", "second", "third"})
	fmt.Println("String queue from slice:", newQueue)

	// Demonstrate common interview use cases
	fmt.Println("\n=== Interview Use Cases ===")

	// 1. Level-order tree traversal simulation
	fmt.Println("1. Level-order Tree Traversal (BFS):")
	levelOrderTraversal()

	// 2. Sliding window maximum
	fmt.Println("\n2. Sliding Window Problems:")
	slidingWindowExample()

	// 3. Generate binary numbers
	fmt.Println("\n3. Generate Binary Numbers:")
	generateBinaryNumbers(5)

	// 4. Hot potato problem (Josephus problem)
	fmt.Println("\n4. Hot Potato Problem:")
	hotPotato([]string{"Alice", "Bob", "Charlie", "David", "Eve"}, 3)
}

// Simulate level-order tree traversal using queue
func levelOrderTraversal() {
	// Simulate a binary tree: 1 -> 2,3 -> 4,5,6,7
	type TreeNode struct {
		val   int
		left  *TreeNode
		right *TreeNode
	}

	// Build tree
	root := &TreeNode{val: 1}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 3}
	root.left.left = &TreeNode{val: 4}
	root.left.right = &TreeNode{val: 5}
	root.right.left = &TreeNode{val: 6}
	root.right.right = &TreeNode{val: 7}

	// BFS traversal using queue
	queue := collections.NewQueue[*TreeNode]()
	queue.Enqueue(root)

	var result []int
	for !queue.IsEmpty() {
		size := queue.Size()
		var level []int

		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			level = append(level, node.val)

			if node.left != nil {
				queue.Enqueue(node.left)
			}
			if node.right != nil {
				queue.Enqueue(node.right)
			}
		}

		result = append(result, level...)
		fmt.Printf("Level: %v\n", level)
	}

	fmt.Printf("Complete traversal: %v\n", result)
}

// Sliding window example using queue
func slidingWindowExample() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Printf("Array: %v, Window size: %d\n", arr, k)

	queue := collections.NewQueue[int]()

	// Process first window
	for i := 0; i < k; i++ {
		queue.Enqueue(arr[i])
	}

	fmt.Printf("Window 1: %v, Sum: %d\n", queue.ToSlice(), sum(queue.ToSlice()))

	// Slide the window
	for i := k; i < len(arr); i++ {
		queue.Dequeue()       // Remove first element
		queue.Enqueue(arr[i]) // Add new element

		fmt.Printf("Window %d: %v, Sum: %d\n", i-k+2, queue.ToSlice(), sum(queue.ToSlice()))
	}
}

// Generate first n binary numbers using queue
func generateBinaryNumbers(n int) {
	if n <= 0 {
		return
	}

	queue := collections.NewQueue[string]()
	queue.Enqueue("1")

	fmt.Printf("First %d binary numbers:\n", n)

	for i := 0; i < n; i++ {
		binary, _ := queue.Dequeue()
		fmt.Printf("%d: %s\n", i+1, binary)

		// Generate next binary numbers by appending 0 and 1
		queue.Enqueue(binary + "0")
		queue.Enqueue(binary + "1")
	}
}

// Hot potato problem (Josephus problem) using queue
func hotPotato(people []string, step int) {
	if len(people) == 0 || step <= 0 {
		return
	}

	queue := collections.NewQueue[string]()
	for _, person := range people {
		queue.Enqueue(person)
	}

	fmt.Printf("Playing hot potato with: %v, step: %d\n", people, step)

	round := 1
	for queue.Size() > 1 {
		// Pass the potato step-1 times
		for i := 0; i < step-1; i++ {
			person, _ := queue.Dequeue()
			queue.Enqueue(person)
		}

		// Remove the person holding the potato
		eliminated, _ := queue.Dequeue()
		fmt.Printf("Round %d: %s is eliminated\n", round, eliminated)
		round++
	}

	winner, _ := queue.Front()
	fmt.Printf("Winner: %s\n", winner)
}

// Helper function to calculate sum
func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
