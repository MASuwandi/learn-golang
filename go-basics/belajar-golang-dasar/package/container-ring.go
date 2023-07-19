package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

// # Package Container/Ring

func main() {
	fmt.Println("# Package Container/Ring: ")

	/*	# Ring
		1 -> 2 -> 3
		⬆		 ⬇
		5 -> - -> 4
	*/
	var data *ring.Ring = ring.New(5)

	for e := 0; e < data.Len(); e++ {
		// data.Value = "Value " + strconv.FormatInt(int64(e), 10)
		data.Value = "Value " + strconv.Itoa(e)
		// or
		// strconv.FormatInt(int64(i), 10)

		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println("value: ", value)
	})

	data.Do(func(a any) {
		fmt.Println("a: ", a)
	})
}
 