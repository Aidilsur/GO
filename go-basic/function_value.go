package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	bye := getGoodBye

	fmt.Println(bye("Aidil"))
}