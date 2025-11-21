package main

import "fmt"

func main() {

	studentMap := make(map[string]int)

	studentMap["tushar"] = 50
	studentMap["rahul"] = 46
	studentMap["ved"] = 86

	
	 for name, marks := range studentMap {
        fmt.Printf("%s has %d marks\n", name, marks)
    }
}
