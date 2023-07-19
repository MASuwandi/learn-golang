package main

import "fmt"

func main() {
	// variable
	fmt.Println("# Variable: ")
	var name string

	name = "Tasty"
	fmt.Println("name: ", name)

	name = "tasty"
	fmt.Println("name: ", name)

	var age = 29
	fmt.Println("age: ", age)

	var son uint8 = 3
	fmt.Println("son: ", son)

	lang := "english"
	fmt.Println("lang: ", lang)

	var (
		firstname string = "John"
		lastname = "Wick"
	)
	fmt.Println("firstname: ", firstname)
	fmt.Println("lastname: ", lastname)

	// const
	const skill1 = "JS"
	const (
		skill2 = "GO"
		skill3 = "JAVA"
	)
	// no print no complaint
}
