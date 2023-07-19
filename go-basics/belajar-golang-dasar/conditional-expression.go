package main

import "fmt"

func main() {
	fmt.Println("# Conditional Expression: ")

	// # If
	name := "not tasty"

	if (name != "tasty") {
		fmt.Println("condition: ", name == "tasty")
		fmt.Println("name: ", name)
	}

	if name == "tasty" {
		fmt.Println("condition: ", name == "tasty")
		fmt.Println("name: ", name)
	} else {
		fmt.Println("not tasty: ", name)
	}

	if name == "tasty" {
		fmt.Println("hello: tasty")
	} else if name == "eko" {
		fmt.Println("hello: eko")
	}
	
	// short statement
	// variable X if
	if length:= len(name); length > 5 {
		fmt.Println("name too long")
	}


	// # Switch
	months := "6"

	switch months {
	case "6":
		fmt.Println("Hello June")
	default:
		fmt.Println("Months unknown number")
	}

	// short state
	switch length := len(months); length > 2 {
	case true:
		fmt.Println("Correct digit length")
	case false:
		fmt.Println("Incorrect digit length")
	}

	// switch without condition
	length := len(months)
	switch {
	case length > 2:
		fmt.Println("Length too long")
	case length <= 2:
		fmt.Println("Length is good")
	}

}
