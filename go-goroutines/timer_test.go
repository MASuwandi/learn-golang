package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer
func TestTimer(t *testing.T)  {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

// time After
func TestTimerAfter(t *testing.T)  {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <- channel
	fmt.Println(time)
}

// time AfterFunc
func TestTimerAfterFunc(t *testing.T)  {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func() {
		fmt.Println("In AfterFunc: ", time.Now())
		group.Done()
	})

	fmt.Println("After AfterFunc: ",time.Now())

	group.Wait()
}
