package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer

	address2.City = "Bandung Barat"
	
	fmt.Println(address1) 
	fmt.Println(address2) 

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) 
	fmt.Println(address2) 
}