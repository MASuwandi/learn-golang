package main

import "fmt"

type Man struct {
	Name	string
}

// without pointer
// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// with pointer
// func (create struct method) funcName() {}
// func ([value from var/this] [struct]) funcName() {}
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fmt.Println("# Pointer In Method")

	// # Pointer In Method
	
	eko := Man{ "Eko" }

	eko.Married()

	fmt.Println("eko.Name: ", eko.Name)
}
 