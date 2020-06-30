package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.4))
	greeting("Howdy")

}

func greeting(name string) {
	fmt.Println(name)
}
