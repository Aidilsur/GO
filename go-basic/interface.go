package main

import "fmt"

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{"Aidil"}
	SayHello(person)
}