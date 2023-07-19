package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld()  {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T)  {
	// async
	go RunHelloWorld()

	fmt.Println("Ups")

	// wait 1 sec before end program
	time.Sleep(1 * time.Second)
}

// Many Goroutine
func DisplayNumber(number int) {
	println("Display ", number)
}

func TestManyGoroutine(t *testing.T)  {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(2 * time.Second)
}