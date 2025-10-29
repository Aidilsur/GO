package main

import "fmt"

func main() {
	name := "asd"

	if name == "Aidil" {
		fmt.Println("Hello Aidil")
		
	} else if name == "Surya" {
		fmt.Println("Hello Surya") 
	} else {
		fmt.Println("Salah") 
	}

   if length := len(name); length > 7 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
	   fmt.Println("Nama Sudah Benar")
   }
}