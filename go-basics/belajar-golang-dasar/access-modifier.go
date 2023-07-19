package main

import (
	"fmt"
	"belajar-golang-dasar/helper"
)

func main() {
	fmt.Println("# Access Modifier: ")

	// # Access Modifier

	// error
	// var help = helper.sayGoodBye("Eko")
	// var ver = helper.version

	var app = helper.Application

	fmt.Println("app: ", app)

	helper.SayHello("Rocket")
}
 