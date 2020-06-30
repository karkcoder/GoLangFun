package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d Hello World %d\n", x, y)
	} else {
		fmt.Printf("%d%d", y, x)
	}

	color := "aaaa"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not found")
	}
}
