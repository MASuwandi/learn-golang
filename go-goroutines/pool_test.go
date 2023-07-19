package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// pool := sync.Pool{}
	// Default Pool Data
	pool := sync.Pool{
		New: func() any {
			return "Default"
		},
	}

	pool.Put("Rocket")
	pool.Put("Racoon")
	pool.Put("Study")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(900 * time.Millisecond)
			pool.Put(data)
		}()
	}

	time.Sleep(1 * time.Second)
}
