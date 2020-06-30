package main

import "fmt"

func main() {
	//Define map
	emails := make(map[string]string)

	emails["tim"] = "tim.karki@orbitshark.com"
	emails["test"] = "test@test.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["test"])
	delete(emails, "test")
	fmt.Println(emails)
}
