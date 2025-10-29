package main

import "fmt"

type User struct {
	Name, Adress string
	Age          int
}

func main() {
	// var user User
	// user.Name = "Aidil Surya"
	// user.Adress = "Jakarta"
	// user.Age = 28

	// fmt.Println(user)

	// user := User {
	// 	Name: "Aidil",
	// 	Adress: "Jakarta",
	// 	Age: 28,
	// }

	user := User{"Aidil", "Jakarta", 28}
	fmt.Println(user)
}