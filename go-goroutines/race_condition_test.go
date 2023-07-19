package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Race Condition
func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func()  {
			for j := 1; j <= 100; j++ {
				x = x + 1
				/* Race Condition
					x = 10 + 1
					x = 10 + 1
					x = 10 + 1
				*/
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Counter = ", x)
}
