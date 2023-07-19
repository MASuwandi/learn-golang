package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Object
// object rep by struct

type Address struct{
	Street		string
	Country		string
	PostalCode	string
}

type Customer struct{
	FirstName	string
	MiddleName	string
	LastName	string
	Age			int
	Married		bool
	Hobbies		[]string 	// slice
	Addresses	[]Address	// Slice of Struct / Array of Object
}

func TestJSONObject(t *testing.T)  {
	customer := Customer{
		FirstName: 	"Rocket",
		MiddleName: "Racoon",
		LastName: 	"Baby",
		Age:		1,
		Married: 	false,
	}
	
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
