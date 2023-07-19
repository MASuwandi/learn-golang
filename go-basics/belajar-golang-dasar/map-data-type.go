package main

import "fmt"

func main() {
	fmt.Println("# Map: ")

	// map[keyType]valueType{key: value}
	person := map[string]string{
		"name": "tasty",
		"address": "Bandung",
	}

	person["title"] = "programmer"

	fmt.Println("person: ", person)
	fmt.Println("person: ", person["name"])
	fmt.Println("person: ", person["address"])

	// make map
	// Create new map
	var book map[string]string = make(map[string]string)
	book2 := make(map[string]string)

	book["title"] = "Belajar Go Lang"
	book["author"] = "tasty"
	book["ups"] = "wrong input"

	// len
	fmt.Println("len(book): ", len(book))
	fmt.Println("book: ", book)
	fmt.Println("book2: ", book2)

	delete(book, "ups")
	fmt.Println("book: ", book)

}
