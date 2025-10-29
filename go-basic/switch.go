package main

import "fmt"

func main() {
	name := "asdasd"

	switch name {
	case "Aidil":
		fmt.Println("Hello Aidil")
	case "Surya":
		fmt.Println("Hello Surya")
	default :
		fmt.Println("Hello")
	}
	
	switch length := len(name); length > 7 {
	case true: 
		fmt.Println("Nama Terlalu Panjang")
	case false: 
		fmt.Println("Nama Sudah Benar")
	}

	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama terlalu Panjang")
	case panjang > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
	
}