package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter);
	// 	counter++
	// }

	// // cara 2
	// for counter2 := 1; counter2 <= 10; counter2++ {
	// 	fmt.Println("Perulangan counter2", counter2)
	// }

	names := []string{"Aidil", "Surya", "Tes"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println((names[i]))
	// }

	for index, name := range names {
		fmt.Println("index :", index, "=", name)
	}


}