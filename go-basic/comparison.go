package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	var equal bool = a == b
	var notEqual bool = a != b
	var greaterThan bool = a > b
	var lessThan bool = a < b

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(greaterThan)
	fmt.Println(lessThan)

}