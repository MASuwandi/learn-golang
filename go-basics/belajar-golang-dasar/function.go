package main

import "fmt"

// func funcName(paramName paramType) returnType { return data }

func getHello(name string) (string, string, string) {
	return "Hello " + name, "good morning", "!"
}

/*
1. Returning Multiple Value
2. Named Returned Values

func funcName() (returnVar1 returnType1, returnVar2 returnType2) { return data1, data2 }

# Invoker
data1, data2 := funcName()
*/
func getName(firstNameIn string, lastNameIn string) (firstName, lastName string) {
	// assign value to return
	firstName = "new"
	lastName = "stuff"
	return
}

// 3. Variadic Function, Var Arg
// slice parameter
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func sayHelloFilter(name string, filter func(string) string) string {
	nameFiltered := filter(name)
	fmt.Println("Hello " + nameFiltered)
	return "Hello " + nameFiltered
}

func spamFilter(name string) string {
	if name == "balls" {
		return "..."
	} else {
		return name
	}
}

func extra(name string, plus string) string {
	fmt.Println("plus: ", plus)
	return name + " " + plus
}

// type declaration
type Filter func(string, string) string

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name, "rocket")
	fmt.Println("Hello plus " + nameFiltered)
	return "Hello " + nameFiltered
}

func main() {
	fmt.Println("# Function: ")

	// # Function
	result1, result2, _ := getHello("tasty")
	fmt.Println(result1, result2)

	firstName, lastName2 := getName("tasty", "go")
	fmt.Println(firstName, lastName2)

	// 3. Var Arg
	total := sumAll(10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10}
	totalSlice := sumAll(numbers...)
	fmt.Println(totalSlice)

	// Passing function and store as value
	goodbye := getGoodBye
	result := goodbye("go")
	fmt.Println(result)

	// Function as parameter
	result3 := sayHelloFilter("rocket go to balls", spamFilter)
	fmt.Println(result3)

	sayHelloFilter("balls", spamFilter)
	sayHelloFilter("rocket", spamFilter)

	sayHelloWithFilter("go", extra)
}
