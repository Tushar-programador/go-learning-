package main

import (
	"fmt"
)

// Corrected function name
func simplefunction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	return a + b
}

// Uses a naked return (optional but common Go style)
func multiple(a, b int) (result int) {
	result = a * b
	return // Naked return
}

// Logical error fixed: performs subtraction and uses naked return
func subtract(a, b int) (result int) {
	result = a - b // FIX: Changed * to -
	return         // Naked return
}

// Naming convention change (unexported)
func divide(a, b int) (result int, err error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	result = a / b
	return result, nil
}

func main() {
	fmt.Println("hello world")
	simplefunction() // Calls the correctly named function

	var a, b int
	fmt.Print("Enter the first number: ") // Using fmt.Print for cleaner input
	_, _ = fmt.Scanln(&a)
	fmt.Print("Enter the second number: ")
	_, _ = fmt.Scanln(&b)

	// --- Calculations ---

	// Add
	addResult := add(a, b)
	fmt.Println("result of adding =", addResult)

	// Multiple
	multResult := multiple(a, b)
	fmt.Println("result of multiply =", multResult)
    
	// Subtract
	subResult := subtract(a, b)
	fmt.Println("result of subtract =", subResult)
    
	// Divide (with proper error handling)
	divResult, err := divide(a, b)
	if err != nil {
		fmt.Println("Error in division:", err) // Print the error if b was 0
	} else {
		fmt.Println("result of divide =", divResult)
	}
}