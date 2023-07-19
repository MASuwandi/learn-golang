package main

import "fmt"

// Struct

type Customer struct {
	Name, Address	string
	Age				int
	Status			bool
}

func main() {
	
	// recommend
	var cust Customer
	cust.Name = "Go"
	cust.Address= "Machine"
	cust.Age = 29

	fmt.Println(cust)
	fmt.Println(cust.Address)

	// Struct Literals
	// type 1
	var joko1 = Customer{
		Name : "Go",
		Address: "Machine",
		Age : 29,
	}

	fmt.Println(joko1)

	// type 2, safe position
	// recommend
	joko2 := Customer{
		Name : "Go",
		Address: "Machine",
		Age : 29,
	}

	fmt.Println(joko2)

	// type 3
	// error if struct type change

	// joko3 := Customer{
	// 	"Go",
	// 	"Machine",
	// 	29,
	// }

	// fmt.Println(joko3)

}
 