package main

import "fmt"

// Interface Kosong
// ex 1

type Any interface {}

func Ups(i Any) Any {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var kosong Any = Ups(3)

	fmt.Println(kosong)
}
  