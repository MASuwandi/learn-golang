package main

import "fmt"

func main() {
	fmt.Println("# Array: ")

	// Array Tipe 1
	var lemari [2]string
	lemari[0] = "GO"
	lemari[1] = "JAVA"

	fmt.Println("lemari: ", lemari)

	// Array Tipe 2
	var bookshelf = [2]string{
		"GO",
		"JAVA",
	}

	fmt.Println("bookshelf: ", bookshelf)
	fmt.Println("bookshelf length: ", len(bookshelf))

	var rak [2]string
	fmt.Println("rak kosong: ", len(rak))
}
