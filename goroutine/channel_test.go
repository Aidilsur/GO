package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string);
	defer close(channel)

	// channel <- "Aidil"
	// data := <- channel

	// fmt.Println(data)
	// defer close(channel)

	go func() {
		time.Sleep((2 * time.Second))
		channel <- "Aidil Surya"

		fmt.Println("selesai mengirim data ke channel")
	} ()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Aidil Surya"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string);
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	
	time.Sleep(5 * time.Second)
}