package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	aidil := Man{"Aidil"}

	aidil.Married()

	fmt.Println(aidil.Name)
}