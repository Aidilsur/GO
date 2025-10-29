package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Aidil")
	data.PushBack("Surya")
	data.PushBack("dilsur")

	var head *list.Element = data.Front()
	
	fmt.Println("head", head.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}