package main

import (
	"fmt"

	"time"
)

func main()  {
	current_time := time.Now()
	fmt.Println("Current tiem is ",current_time)

	formmatted_time := current_time.Format("2006/01/02/15/04/05")
		fmt.Println("Current tiem is ", formmatted_time )

	formmatted_time= current_time.Format("2005/05/26") // cannpt put any data for pattern
		fmt.Println("Current tiem is ", formmatted_time )

}