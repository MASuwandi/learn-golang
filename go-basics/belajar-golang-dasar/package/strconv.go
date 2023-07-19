package main

import (
	"fmt"
	"strconv"
)

// # Package strconv

func main() {
	fmt.Println("# Package strconv: ")

	// ParseBool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("boolean: ", boolean)
	} else {
		// test without .Error()
		fmt.Println("err.Error(): ", err.Error())
	}

	// // ParseFloat
	// float, err := strconv.ParseFloat()
	// if err == nil {
	// 	fmt.Println("float: ", float)
	// } else {
	// 	fmt.Println("err.Error(): ", err.Error())
	// }

	// ParseInt(sring, base, bit size)
	parseInt, err := strconv.ParseInt("1000", 10, 64)
	if err == nil {
		fmt.Println("parseInt: ", parseInt)
	} else {
		fmt.Println("err.Error(): ", err.Error())
	}

	// // FormatBool
	// formatBool := strconv.FormatBool()
	// if err == nil {
	// 	fmt.Println("formatBool: ", formatBool)
	// } else {
	// 	fmt.Println("err.Error(): ", err.Error())
	// }
	
	// // FormatFloat
	// formatFloat := strconv.FormatFloat()
	// if err == nil {
	// 	fmt.Println("formatFloat: ", formatFloat)
	// } else {
	// 	fmt.Println("err.Error(): ", err.Error())
	// }
	
	// FormatInt(string, base)
	// base: decimal/10, binary/2, octal/8, hexadecimal/16
	formatInt := strconv.FormatInt(1000, 10)
	fmt.Println("formatInt: ", formatInt)
	
	// FormatFloat
	valueInt, err := strconv.Atoi("200")
	if err == nil {
		fmt.Println("valueInt: ", valueInt)
	} else {
		fmt.Println("err.Error(): ", err.Error())
	}
}
 