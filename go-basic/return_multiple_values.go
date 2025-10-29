package main

import "fmt"

func getFullName() (string, string) {
	return "Aidil", "Surya"
}

func main() {
	// firstName, lastName := getFullName()

	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()

	fmt.Println(firstName)
}