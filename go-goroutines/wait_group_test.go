package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup)  {
	defer group.Done()

	group.Add(1)

	fmt.Println("After Add Proses")
	time.Sleep(500 * time.Millisecond)
}

func TestWaitGroup(t *testing.T)  {
	// await group
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("TestWaitGroup Done")
}