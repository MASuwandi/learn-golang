package go_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Gomaxprocs
func TestGetGomaxprocs(t *testing.T)  {
	group := sync.WaitGroup{}

	// Increase created goroutines
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func ()  {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	// Total CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	// Total Thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread: ", totalThread)

	// Running Goroutines
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutines: ", totalGoroutine)

	group.Wait()
}

func TestChangeGomaxprocs(t *testing.T)  {
	group := sync.WaitGroup{}

	// Increase created goroutines
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func ()  {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	// Total Thread
	runtime.GOMAXPROCS(20) // Increase
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread: ", totalThread)

	// Running Goroutines
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutines: ", totalGoroutine)

	group.Wait()
}

/* # Recommend

no need to change Thread,
go lang already optimal with Thread management.

better horizontal scaling
*/
