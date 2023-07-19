package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// JSON Tag

type Product struct{
	Id			string `json:"id"`
	Name		string `json:"name"`
	ImageURL	string `json:"image_url"`
}

func TestTag(t *testing.T)  {
	product := Product{
		Id: 		"P001",
		Name: 		"Baby Racoon",
		ImageURL: 	"http://example.com/image.png",
	}
	
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

// JSON Tag Decode
func TestTagDecode(t *testing.T)  {
	// Case 1
	jsonString := `{"id":"P001","name":"Baby Racoon","image_url":"http://example.com/image.png"}`

	// Case 2
	// jsonString := `{"Id":"P001","Name":"Baby Racoon","ImageURL":"http://example.com/image.png"}`

	// When read not case sensitive
	// ImageURL unreadable cause diffrent char
	jsonBytes := []byte(jsonString)

	product := &Product{}
	
	json.Unmarshal(jsonBytes, product)

	fmt.Println("product: ", product)
}
