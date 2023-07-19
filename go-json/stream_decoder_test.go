package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Stream Decoder
func TestStreamDecoder(t *testing.T)  {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)
	
	customer := Customer{}

	// Like Unmarshal
	decoder.Decode(&customer)

	fmt.Println("customer: ", customer)
}
