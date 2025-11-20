package main

import "fmt"

func main() {
	var arr1 = [2]int{}
	fmt.Println("First array =", arr1)
	arr4 := [6]int{1, 2, 3} 
	fmt.Println("First array =", arr4)
	
	arr2 := []int{1, 2, 3} 
	fmt.Println("First array =", arr2)
	arr3 := [5]int{2: 10, 3: 25} // add valu ad index

	fmt.Println("First array =", arr3)

}
