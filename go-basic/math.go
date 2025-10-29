package main

import "fmt"

func main() {
	var a = 25
	var b = 5

	var plus = a + b
	var substract = a - b
	var multiply = a * b
	var divide = a / b
	var modulus = a % b

	var i = 10
	i += 10 // i = i + 10

	var j = 1
	var k = 2
	j++ // j = j + 1
	k-- // j = j + 1

	fmt.Println(plus)
	fmt.Println(substract)
	fmt.Println(multiply)
	fmt.Println(divide)
	fmt.Println(modulus)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}