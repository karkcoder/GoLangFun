package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}