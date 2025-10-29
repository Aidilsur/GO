package main

import "fmt"

func main() {
	firstName := "Aidil"
	lastName := "Surya"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'", firstName, lastName)
}