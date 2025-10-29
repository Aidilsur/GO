package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Aidil"
	lastName = "Surya"

	return firstName, lastName
} 

func main() {
	a, b := getCompleteName()

	fmt.Println(a, b)

}