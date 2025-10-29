package main

import "fmt"

type Filter func(string) string

func sayHelloWithfilter(name string, filter Filter) {
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