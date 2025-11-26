package main

import "fmt"

func main () {
	var num *int
	number := 5
	num = &number

	fmt.Println(num)
}