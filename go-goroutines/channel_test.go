package go_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_CreateChannel(t *testing.T) {
	channel := make(chan string)

	// close channel
	defer close(channel)

	// send data
	channel <- "Rocket"

	// receive data
	receiver := <-channel

	fmt.Println(<-channel)
	fmt.Println(receiver)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Rocket Racoon"

		fmt.Println("Sending Data to Channel Done")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

// Channel as Parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Rocket Racoon"

	fmt.Println("Sending Data to Channel Done")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

// Channel In dan Out
func OnlyIn(channel chan<- string) {
	channel <- "Rocket Racoon"

	fmt.Println("Sending Data to Channel Done")
}

func OnlyOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)
	fmt.Println("Receive Data from Channel Done")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

// Buffered Channel
// Buffer not blocking
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Rocket"
	channel <- "Racoon"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	fmt.Println("Buffered Channel Done")
}

func TestBufferedChannelGoRou(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Rocket"
		channel <- "Racoon"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Buffered Channel Done")
}

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Rocket: " + strconv.Itoa(i)
		}
		// close so not deadlock
		close(channel)
	}()

	// if not close, will loop forever
	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Range Channel Done")
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	nChannel := 0
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		}

		// stop channel from deadlock
		if counter == nChannel {
			break
		}
	}

	fmt.Println("Select Channel Done")
}

// Default Channel
func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	nChannel := 0
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		default:
			fmt.Println("Waiting for Data from Channel: ")
		}
		
		if counter == nChannel {
			break
		}
	}

	fmt.Println("Select Channel Done")
}
