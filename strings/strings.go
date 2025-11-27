package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	num = 9
	fmt.Println("Number is : ",num);
	fmt.Printf("Numebr type is %T\n",num);
	
	var Number_to_String string =strconv.Itoa(num);
	fmt.Println("Number is : ",Number_to_String);
	fmt.Printf("Numebr type is %T\n",Number_to_String);

	var string_to_Number float64 ;
	string_to_Number ,_  = strconv.ParseFloat(Number_to_String , 64)

	fmt.Println("Number is : ",string_to_Number);
	fmt.Printf("Numebr type is %T\n",string_to_Number);


}