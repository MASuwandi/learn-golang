package main

import "fmt"

func main() {
	fmt.Println("# Conversion: ")

	// Conversion Number
	var nilai32 uint32 = 128
	fmt.Println("nilai32: ", nilai32)

	var nilai8 uint8 = uint8(nilai32)
	fmt.Println("nilai8: ", nilai8)

	// Conversion Byte to String
	var name = "Eko Khan"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println("name: ", name)
	fmt.Println("e: ", e)
	fmt.Println("eString: ", eString)

}
