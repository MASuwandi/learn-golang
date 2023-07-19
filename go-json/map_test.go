package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Map
func TestMapDecode(t *testing.T)  {
	jsonString := `{
		"id": "P001",
		"name": "Baby Racoon",
		"price": 20000000
	}`

	// When read not case sensitive
	// ImageURL unreadable cause diffrent char
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println("result: ", result)
	fmt.Println("result id: ", result["id"])
	fmt.Println("result name: ", result["name"])
	fmt.Println("result price: ", result["price"])
}

func TestMapEncode(t *testing.T)  {
	product := map[string]interface{}{
		"id": "P001",
		"name": "Baby Racoon",
		"price": 20000000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes: ", string(bytes))
}
