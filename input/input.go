package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	fmt.Println("hello guys , What is your name ?")
	var name string;

	// fmt.Scan(&name)

	// fmt.Println("my name is", name)

	reader := bufio.NewReader(os.Stdin)
	name,_ = reader.ReadString('\n')


	fmt.Println("my name is", name)
}