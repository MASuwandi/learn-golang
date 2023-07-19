package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Create Context
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println("background: ", background)

	todo := context.TODO()
	fmt.Println("todo: ", todo)
}

// Context with value
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// add data = create child
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")

	fmt.Println("contextA: ", contextA)
	fmt.Println("contextB: ", contextB)
	fmt.Println("contextC: ", contextC)
	fmt.Println("contextD: ", contextD)
	fmt.Println("contextE: ", contextE)
	fmt.Println("contextF: ", contextF)
	fmt.Println("contextG: ", contextG)

	// Get Context Value
	fmt.Println("context F value f: ", contextF.Value("f"))
	fmt.Println("context F value c: ", contextF.Value("c"))
	fmt.Println("context F value b: ", contextF.Value("b"))
	fmt.Println("context A value b: ", contextA.Value("b"))
}

// Context with cancel

// goroutine leak scenario
// func CreateCounter() chan int {
// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			destination <- counter
// 			counter++
// 		}
// 	}()

// 	return destination
// }

// func TestContextWithoutCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

// 	destination := CreateCounter()

// 	for n := range destination {
// 		fmt.Println("Counter: ", n)
// 		if n == 10 {
// 			break
// 		}
// 	}

// 	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
// }

// goroutine without leak scenario
func CreateCounter(ctx context.Context) chan int {
	// ctx listen to context signal
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine - start: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine before sleep - after create goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}

	cancel() // send signal cancel to context

	fmt.Println("Total Goroutine - before sleep: ", runtime.NumGoroutine())

	time.Sleep(1 * time.Second)

	fmt.Println("Total Goroutine - after sleep: ", runtime.NumGoroutine())
}

/* # Website case
one request one goroutine leak
1000 request 1000 goroutine leak
*/

/*
Recommend for cancel DB query using cancel
*/


// Context With Timeout

// Scenario 1 - with sleep
func CounterTimeout(ctx context.Context) chan int {
	// ctx listen to context signal
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine - start: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	defer cancel()

	destination := CounterTimeout(ctx)

	fmt.Println("Total Goroutine before sleep - after create goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}


// Context With Deadline
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine - start: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	defer cancel()

	destination := CounterTimeout(ctx)

	fmt.Println("Total Goroutine before sleep - after create goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}
