package main

import "fmt"

func main() {
	fmt.Println("# Type Declaration: ")

	// Type Declaration
	type NoKTP string
	type Married bool

	var nik NoKTP = "1x9"
	var status Married = true

	fmt.Println("nik: ", nik)
	fmt.Println("status: ", status)
}
