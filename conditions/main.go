package main

import "fmt"

func main() {
	num := 20

	if num > 5 {
		fmt.Println("Number is Greater than 5:", num)
	} else {
		fmt.Println("Number is not greater than 5:", num)
	}

	switch num {
	case 10:
		fmt.Println("Number is 10");
	case 20: 
		fmt.Println("Number is 20");
	
	default:
		fmt.Println("Number is",num)
	
	}
}