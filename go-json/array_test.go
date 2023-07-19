package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Array
// array rep by slice

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Rocket",
		MiddleName: "Racoon",
		LastName:   "Baby",
		Age:        1,
		Married:    false,
		Hobbies:    []string{"Code", "Eat", "Sleep"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
	fmt.Println(customer)
}

func TestArrayDecode(t *testing.T) {
	jsonString := `{
		"FirstName": 	"Rocket",
		"MiddleName": 	"Racoon",
		"LastName": 	"Baby",
		"Age":			1,
		"Married": 		false,
		"Hobbies": 		["Code", "Eat", "Sleep"]
	}`
	jsonBytes := []byte(jsonString)

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
	fmt.Println("customer Hobbies: ", customer.Hobbies)
}

func TestArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Rocket",
		Addresses: []Address{
			{
				Street:     "Wakanda St",
				Country:    "Wakanda",
				PostalCode: "333",
			},
			{
				Street:     "Asgard St",
				Country:    "Asgard",
				PostalCode: "888",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes: ", string(bytes))
}

func TestArrayComplexDecode(t *testing.T) {
	jsonString := `{
		"FirstName": 	"Rocket",
		"Addresses":	[
			{
				"Street":"Wakanda St","Country":"Wakanda","PostalCode":"333"
			},
			{
				"Street":"Asgard St","Country":"Asgard","PostalCode":"888"
			}
		]
	}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	// send pointer so the change will affect the customer not the copy.
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println("customer: ", customer)
	fmt.Println("customer FirstName: ", customer.FirstName)
	fmt.Println("customer Addresses: ", customer.Addresses)
}

// Decode JSON Array
func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[
			{
				"Street":"Wakanda St","Country":"Wakanda","PostalCode":"333"
			},
			{
				"Street":"Asgard St","Country":"Asgard","PostalCode":"888"
			}
		]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	// send pointer so the change will affect the customer not the copy.
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println("addresses: ", addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Wakanda St",
			Country:    "Wakanda",
			PostalCode: "333",
		},
		{
			Street:     "Asgard St",
			Country:    "Asgard",
			PostalCode: "888",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes: ", string(bytes))
}
