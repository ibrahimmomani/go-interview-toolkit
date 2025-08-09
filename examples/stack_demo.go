package main

import (
	"fmt"
	"log"

	"github.com/ibrahimmomani/go-interview-toolkit/collections"
)

func main() {
	fmt.Println("=== Stack Demo ===")

	// Create a new stack
	stack := collections.NewStack[int]()
	fmt.Println("Created empty stack:", stack)

	// Push elements (LIFO - Last In, First Out)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("After pushing 1, 2, 3:", stack)

	// Peek at top element
	if top, err := stack.Peek(); err == nil {
		fmt.Printf("Top element (peek): %d\n", top)
	}

	// Pop elements
	if val, err := stack.Pop(); err == nil {
		fmt.Printf("Popped: %d, Stack now: %s\n", val, stack)
	}

	// Multi-push
	stack.MultiPush(10, 20, 30)
	fmt.Println("After multi-push 10, 20, 30:", stack)

	// Multi-pop
	if values, err := stack.MultiPop(2); err == nil {
		fmt.Printf("Multi-popped: %v, Stack now: %s\n", values, stack)
	}

	// Peek multiple elements
	if values, err := stack.PeekN(3); err == nil {
		fmt.Printf("Top 3 elements (peek): %v\n", values)
	}

	// Convert to slice
	slice := stack.ToSlice()
	fmt.Printf("As slice: %v\n", slice)

	// Clone the stack
	clone := stack.Clone()
	fmt.Println("Cloned stack:", clone)

	// Reverse the original
	stack.Reverse()
	fmt.Println("After reversing original:", stack)
	fmt.Println("Clone remains unchanged:", clone)

	// Create from slice
	newStack := collections.FromSliceStack([]string{"bottom", "middle", "top"})
	fmt.Println("String stack from slice:", newStack)

	// Demonstrate common interview use cases
	fmt.Println("\n=== Interview Use Cases ===")

	// 1. Balanced Parentheses
	fmt.Println("1. Balanced Parentheses Check:")
	expressions := []string{"()", "([{}])", "({[})", "((()))", "({[}]"}
	for _, expr := range expressions {
		balanced := isBalanced(expr)
		fmt.Printf("'%s' is balanced: %t\n", expr, balanced)
	}

	// 2. Reverse a string using stack
	fmt.Println("\n2. Reverse String:")
	original := "Hello World"
	reversed := reverseString(original)
	fmt.Printf("'%s' reversed: '%s'\n", original, reversed)

	// 3. Evaluate postfix expression
	fmt.Println("\n3. Postfix Expression Evaluation:")
	postfix := []string{"3", "4", "+", "2", "*", "1", "-"}
	result := evaluatePostfix(postfix)
	fmt.Printf("Postfix %v = %d\n", postfix, result)
}

// Common interview problem: Check if parentheses are balanced
func isBalanced(s string) bool {
	stack := collections.NewStack[rune]()
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack.Push(char)
		case ')', ']', '}':
			if stack.IsEmpty() {
				return false
			}
			if top, err := stack.Pop(); err != nil || top != pairs[char] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

// Reverse a string using stack
func reverseString(s string) string {
	stack := collections.NewStack[rune]()

	// Push all characters
	for _, char := range s {
		stack.Push(char)
	}

	// Pop all characters to build reversed string
	var result []rune
	for !stack.IsEmpty() {
		if char, err := stack.Pop(); err == nil {
			result = append(result, char)
		}
	}

	return string(result)
}

// Evaluate postfix expression
func evaluatePostfix(tokens []string) int {
	stack := collections.NewStack[int]()

	for _, token := range tokens {
		switch token {
		case "+":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a + b)
		case "-":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a - b)
		case "*":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a * b)
		case "/":
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			if b != 0 {
				stack.Push(a / b)
			} else {
				log.Fatal("Division by zero")
			}
		default:
			// It's a number
			var num int
			fmt.Sscanf(token, "%d", &num)
			stack.Push(num)
		}
	}

	result, _ := stack.Pop()
	return result
}
