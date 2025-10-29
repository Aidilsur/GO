package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "my name is ", customer.Name)
}

func main() {
	user := Customer{"Aidil", "Jakarta", 28}
	fmt.Println(user)

	user.sayHello("AI")
}