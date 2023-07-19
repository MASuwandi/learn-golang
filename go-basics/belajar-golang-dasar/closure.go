package main

import "fmt"

// # Closure Function
// update from different scope

func main() {
	counter := 0
	increment := func ()  {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
