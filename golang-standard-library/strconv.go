package main

import (
	"fmt"
	"strconv"
)

func main() {
	res, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(res)
	}

	resInt, errInt := strconv.Atoi("1000")
	if errInt != nil {
		fmt.Println("Error", errInt.Error())
	} else {
		fmt.Println(resInt)
	}

	bin := strconv.FormatInt(999, 2)
	fmt.Println(bin)
	
	resItoa := strconv.Itoa(999)
	fmt.Println(resItoa)
}