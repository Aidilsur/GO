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

// channel as paramater
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

// channel in & out

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Aidil Surya"
}

func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOuChannel(t *testing.T) {
	channel := make(chan string);
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)


	time.Sleep(5 * time.Second)
}

// Channel Buffer
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Aidil"
	channel <- "Surya"

	fmt.Println(<- channel)
	fmt.Println(<- channel)

	fmt.Println("selesai")
}

func TestBufferedChannel2(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Aidil"
		channel <- "Surya"
	} ()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	} ()

	time.Sleep(2 * time.Second)


	fmt.Println("selesai")
}