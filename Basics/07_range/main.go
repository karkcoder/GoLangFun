package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	ids := []int{11, 22, 33, 4, 5, 64}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}
}
