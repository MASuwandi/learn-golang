package helper

import "fmt"

// # Package

func SayHello(name string) {
	fmt.Println("Hello " + name)
}

// # Access Modifier
var version = "1.0.0" // tidak bisa diakses dari luar
var Application = "Belajar Golang"

// Tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Good Bye " + name
}
 