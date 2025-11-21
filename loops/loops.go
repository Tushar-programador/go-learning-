package main

import "fmt"

func main()  {
	
	counter:=0

	for i := 0; i < 10; i++ {
		counter++;
		println("Counter",counter)
	}

	for range 10 {
		counter++;
		println("Counter",counter)
	}
	arr :=[10]int{5,6,8,9,4,5,5,5,6}
	for index ,value:= range(arr){
		fmt.Printf("Value at %d index is %d\n",index , value)
	}

	

}
