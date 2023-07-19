package main

import (
	"fmt"
	"math"
)

// # Package math

func main() {
	fmt.Println("# Package math: ")

	// Round
	round1 := math.Round(1.5) // upper limit
	round2 := math.Round(1.4) // lower limit

	// Floor
	floor := math.Floor(1.9)

	// Ceil
	ceil := math.Ceil(1.1)

	// Max
	max := math.Max(10, 20)

	// Min
	min := math.Min(10, 20)

	fmt.Println("round1: ", round1)
	fmt.Println("round2: ", round2)
	fmt.Println("floor: ", floor)
	fmt.Println("ceil: ", ceil)
	fmt.Println("max: ", max)
	fmt.Println("min: ", min)
}
 