package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Decode JSON

func TestDecode(t *testing.T)  {
	jsonRequest := `{
		"FirstName": 	"Rocket",
		"MiddleName": "Racoon",
		"LastName": 	"Baby",
		"Age":		1,
		"Married": 	false
	}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}

	// send pointer so the change will affect the customer not the copy.
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("customer: ", customer)
	fmt.Println("customer FirstName: ", customer.FirstName)
	fmt.Println("customer Age: ", customer.Age)
	fmt.Println("customer Married: ", customer.Married)
}
