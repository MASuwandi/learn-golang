package main

import "fmt"

// Struct

type Customer struct {
	Name, Address 	string
	Age				int
	Status			bool
}

// Set func as struct method Customer
func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is ", customer.Name)
}

func main() {
	
	rully := Customer{
		Name : "Go",
	}
	rully.sayHello()

	fmt.Println(rully)
}
 