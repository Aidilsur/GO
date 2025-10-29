package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Aidil"
	names[1] = "Surya"

	fmt.Println(names);
	fmt.Println(names[0]);
	fmt.Println(names[1]);
	
	
	var values = [3]int {
		90, 80, 95,
	}
	
	fmt.Println(values);
	fmt.Println(len(values))

}