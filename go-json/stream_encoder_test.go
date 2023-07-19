package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Stream Encoder
func TestStreamEncoder(t *testing.T)  {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)
	
	customer := Customer{
		FirstName: 	"Rocket",
		LastName: 	"Racoon",
	}

	encoder.Encode(customer)

	fmt.Println("customer: ", customer)
}
