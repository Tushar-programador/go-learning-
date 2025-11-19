package main

import (
	"fmt"
)

func simplefucntion() {
	println("simpel fucntion")

}
func add(a, b int) int {
	return a + b
}
func multiple(a, b int) (result int) {

	result = a * b

	return result

}
func subtract(a, b int) (result int) {

	result = a * b

	return result

}

func main() {

	fmt.Printf("hello world")
	simplefucntion()
	var a, b int
	println("Enter the first number")
	_, _ = fmt.Scanln(&a)
	println("Enter the second number")
	_, _ = fmt.Scanln(&b)

	result := add(a, b)
	println("reslut of adding 3 + 5 =", result)

	result = multiple(a, b)
	println("reslut of multiplr 3 + 5 =", result)
	result = subtract(a, b)
	println("reslut of subtract 3 + 5 =", result)

}
