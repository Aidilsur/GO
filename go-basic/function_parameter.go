package main

import "fmt"

func sayHelloWithfilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	str := "anjing"
	
	sayHelloWithfilter(str, spamFilter)

}