package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Aidil",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	
	book := map[string]string{
		"tittle": "Buku Golang",
		"author": "Aidil Surya",
		"ups": "Salah",
	}
	
	fmt.Println(book)
	
	delete(book, "ups")
	
	fmt.Println(book)
}