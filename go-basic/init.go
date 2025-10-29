package main

import (
	"fmt"
	"go-basics-1/database"
	_ "go-basics-1/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}