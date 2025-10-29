package main

import (
	"fmt"
	"go-basics-1/helper"
)

func main() {
	result := helper.SayHello("Aidil")

	fmt.Println(result)
	
	fmt.Println(helper.Application)
	fmt.Println(helper.version) // tidak bisa di akses
	fmt.Println(helper.sayGoodBye("Aidil")) // tidak bisa di akses
}