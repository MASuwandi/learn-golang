package main

import "fmt"

func main() {
	fmt.Println("# Slice: ")

	// pointer, capacity, length
	// Slice Tipe 1

	// Array
	var months = [...]string{
		"January", "February",
		"March", "April",
		"May", "June",
		"July", "August",
		"September", "October",
		"November", "December",
	}

	fmt.Println("months: ", months)

	// Slice
	slice1 := months[4:7]
	/*
	pointer 4	-> index 4, start data
	length 3	-> 4, 5, 6
	capacity 8 	-> from pointer to end
	*/

	fmt.Println("slice1: ", slice1)

	slice1[0] = "May Update"
	fmt.Println("slice1: ", slice1)

	fmt.Println("slice1: ", len(slice1))
	fmt.Println("slice1: ", cap(slice1))
	fmt.Println("months: ", months)

	// Append
	// create new array if cap full
	append1 := append(slice1, "Bulan Baru")
	fmt.Println("append1: ", append1)

	// Make
	// create slice from 0
	// ([]type, length, capacity)

	makeSlice1 := make([]string, 2, 5)
	makeSlice1[0] = "Eko"
	makeSlice1[1] = "Khan"

	fmt.Println("makeSlice1: ", makeSlice1)
	fmt.Println("len: ", len(makeSlice1))
	fmt.Println("cap: ", cap(makeSlice1))

	// Copy
	// copy slice
	copySlice1 := make([]string, 2, 2)

	copy(copySlice1, makeSlice1)
	fmt.Println("copySlice1: ", copySlice1)

	arr1 := []int{1,2}
	arr2 := []int{3,4}

	copy1 := copy(arr1, arr2)

	fmt.Println("arr1: ", arr1)
	fmt.Println("arr2: ", arr2)
	fmt.Println("copy1: ", copy1)

	// this is array
	iniArray := [...]int{5,6}

	// this is slice
	iniSlice := []int{7,8}

	fmt.Println("iniArray: ", iniArray)
	fmt.Println("iniSlice: ", iniSlice)
}
