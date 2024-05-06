package main

import (
	"fmt"
)

// @inline
// Adds two integers and returns the sum
func add(a int, b int) int {
	return a + b
}

// @inline
// Multiplies two integers and returns the result
func multiply(a int, b int) int {
	return a * b
}

// Uses the add function in a loop
func sumLoop(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total = add(total, num)
	}
	return total
}

// Uses the multiply function in a conditional
func conditionalProduct(a int, b int) int {
	if a > 0 {
		return multiply(a, b)
	}
	return 0
}

// Main function to demonstrate function usage
func atry() {
	fmt.Println("Adding 3 and 4:", add(3, 4))
	fmt.Println("Multiplying 3 and 4:", multiply(3, 4))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of numbers:", sumLoop(numbers))
	fmt.Println("Conditional product:", conditionalProduct(3, 100))
}
