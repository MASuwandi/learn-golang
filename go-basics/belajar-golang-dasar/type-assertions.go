package main

import (
	"fmt"
)

// # Type Assertions

func random() interface{} {
	return 123
}

func main() {

	result := random()
	// # Set data type to string
	// resultString := result.(string)
	// fmt.Println("resultString: ", resultString)

	// # Set data type to Number
	// resultInt := result.(int)
	// fmt.Println("resultInt: ", resultInt)

	// # Set data type to Boolean
	// resultBool := result.(bool)
	// fmt.Println("resultInt: ", resultBool)

	switch value := result.(type) {
	case string:
		fmt.Println("String: ", value)
	case int:
		fmt.Println("Int: ", value)
	default:
		fmt.Println("Unknown: ")
	}
}
 