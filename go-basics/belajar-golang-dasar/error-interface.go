package main

import (
	"fmt"
	"errors"
)

// # Error Interface

// built in
// type error interface {
// 	Error()	string
// }

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// var contohError error = errors.New("Ups Error")
	// fmt.Println(contohError)

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("err: ", err)
		fmt.Println("err.Error(): ", err.Error())
	}
}
 