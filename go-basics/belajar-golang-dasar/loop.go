package main

import "fmt"

func main() {
	fmt.Println("# Loop: ")

	// # Looping
	counter := 0

	for counter <= 10 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	for counter := 0; counter <= 10; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	// data collection array, slice, map

	// slice
	slices := []string{"a", "b"}

	// manual
	for i := 0; i < len(slices); i++ {
		fmt.Println("Perulangan ke: ", slices[i])
	}

	// for with range
	for i, slice :=  range slices {
		fmt.Println("Perulangan ke: ", i, slice)
	}

	for _, slice :=  range slices {
		fmt.Println("Perulangan ke: ", slice)
	}

	// map
	person := make(map[string]string)
	person["name"] = "tasty"
	person["title"] = "programmer"

	for key, value :=  range person {
		fmt.Println("Perulangan ke: ",key, value)
	}


	// # Break & Continue
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke: ",i)
		if i == 5 {
			break
		}
	}
}
 