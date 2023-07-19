package main

import "fmt"

// # Nil
// can be use by interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}
func main() {
	// ver 1
	// var person1 = map[string]string{"name": ""}
	var person1 map[string]string

	if person1["name"] == "" {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person1)
	}

	// ver 2
	var person2 map[string]string = NewMap("")
	if person2 == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person2)
	}

	var person3 map[string]string = NewMap("Rocket")
	if person3 == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person3)
	}
}
 