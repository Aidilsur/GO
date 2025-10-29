package goroutine

import (
	"fmt"
	"strconv"
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

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func () {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	} ()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data:= <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++

		case data:= <- channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		} 
	}
}