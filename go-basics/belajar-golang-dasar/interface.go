package main

import "fmt"

// Interface
// ex 1

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

type Person struct {
	Name string
}

// without this var base on Person would error
// make func as struct method
func (person Person) GetName() string {
	return person.Name
}

// ex 2
type Animal struct {
	Name 	string
	Age 	uint8
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	animal := Animal{ Name: "Cambo" }
	SayHello(animal)
}
 