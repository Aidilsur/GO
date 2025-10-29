package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Aidil Surya"

	fmt.Println(strings.Contains(str, "Aidil"))
	fmt.Println(strings.Split(str, " "))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Trim("   Aidil Surya   ", " "))
	fmt.Println(strings.ReplaceAll(str, "Aidil", "Aidils"))
}