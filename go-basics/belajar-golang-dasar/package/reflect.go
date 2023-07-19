package main

import (
	"fmt"
	"reflect"
)

// # Package Reflect

// type Sample struct {
// 	Name	string
// }

// StructTag
type Sample struct {
	Name string `required:"true" max:"10"`
}

// validation library example
type Example struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func isValid(data any) bool {
	typee := reflect.TypeOf(data)
	for i := 0; i < typee.NumField(); i++ {
		field := typee.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
			// or
			// value := reflect.ValueOf(data).Field(i).Interface()
			// if value == "" {
			// 	return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("# Package Reflect: ")

	sample := Sample{"Eko"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println("sampleType.NumField: ", sampleType.NumField())

	structField := sampleType.Field(0)

	fmt.Println("structField.Name: ", structField.Name)

	// StructTag
	required := structField.Tag.Get("required")
	max := structField.Tag.Get("max")

	fmt.Println("required: ", required)
	fmt.Println("max: ", max)

	// invoke validation library
	sample.Name = ""
	fmt.Println("isValid sample: ", isValid(sample))

	ex := Example{"Eko", "test"}
	fmt.Println("isValid ex: ", isValid(ex))

}
