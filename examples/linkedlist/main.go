package main

import (
	"fmt"
	"log"

	"github.com/ibrahimmomani/go-interview-toolkit/collections"
)

func main() {
	// Create a new linked list
	ll := collections.NewLinkedList[int]()
	fmt.Println("Created empty list:", ll)

	// Add elements
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	fmt.Println("After appending 1, 2, 3:", ll)

	// Prepend element
	ll.Prepend(0)
	fmt.Println("After prepending 0:", ll)

	// Insert at specific position
	if err := ll.Insert(2, 15); err != nil {
		log.Fatal(err)
	}
	fmt.Println("After inserting 15 at index 2:", ll)

	// Get element at index
	if value, err := ll.Get(2); err == nil {
		fmt.Printf("Element at index 2: %d\n", value)
	}

	// Find element
	index := ll.Find(2)
	fmt.Printf("Element 2 found at index: %d\n", index)

	// Check if contains
	fmt.Printf("Contains 15: %t\n", ll.Contains(15))
	fmt.Printf("Contains 99: %t\n", ll.Contains(99))

	// Get head and tail
	if head, err := ll.Head(); err == nil {
		fmt.Printf("Head: %d\n", head)
	}
	if tail, err := ll.Tail(); err == nil {
		fmt.Printf("Tail: %d\n", tail)
	}

	// Convert to slice
	slice := ll.ToSlice()
	fmt.Printf("As slice: %v\n", slice)

	// Delete element
	deleted := ll.Delete(15)
	fmt.Printf("Deleted 15: %t, List: %s\n", deleted, ll)

	// Delete at index
	if err := ll.DeleteAt(1); err == nil {
		fmt.Println("After deleting at index 1:", ll)
	}

	// Reverse the list
	ll.Reverse()
	fmt.Println("After reversing:", ll)

	// Size and empty check
	fmt.Printf("Size: %d, Is empty: %t\n", ll.Size(), ll.IsEmpty())

	// Create from slice
	newLL := collections.FromSlice([]string{"hello", "world", "!"})
	fmt.Println("String list from slice:", newLL)

	// Clear
	ll.Clear()
	fmt.Printf("After clearing - Size: %d, Is empty: %t\n", ll.Size(), ll.IsEmpty())

	// Demonstrate with different types
	floatList := collections.NewLinkedList[float64]()
	floatList.Append(3.14)
	floatList.Append(2.71)
	fmt.Println("Float list:", floatList)

	// Custom struct example
	type Person struct {
		Name string
		Age  int
	}

	personList := collections.NewLinkedList[Person]()
	personList.Append(Person{Name: "Alice", Age: 30})
	personList.Append(Person{Name: "Bob", Age: 25})
	fmt.Println("Person list:", personList)
}
