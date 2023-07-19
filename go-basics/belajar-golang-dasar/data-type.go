package main

import "fmt"

func main() {
	// number
	fmt.Println("# Number Data Type: ")
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma dua = ", 3.2)

	// boolean
	fmt.Println("\n# Boolean Data Type: ")
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// string
	fmt.Println("\n# String Data Type: ")
	fmt.Println("Tasty")
	fmt.Println("GO to the moon")

	// Function for
	fmt.Println("\n# String Function: ")
	fmt.Println("\n# Length Method: ")
	fmt.Println(len("Tasty"))
	fmt.Println("\n# Accessing Index: ")
	fmt.Println("Tasty"[0])
}
