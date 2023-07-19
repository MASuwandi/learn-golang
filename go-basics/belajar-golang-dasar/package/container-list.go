package main

import (
	"container/list"
	"fmt"
)

// # Package Container/List

func main() {
	fmt.Println("# Package Container/List: ")

	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Khandy")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("e: ", e)
		fmt.Println("e.Value: ", e.Value)
	}

	fmt.Println("data.Front() - 22: ", data.Front().Value)
	fmt.Println("data.Back() - 23: ", data.Back().Value)
	fmt.Println("data.Front() 24: ", data.Back().Prev().Value)

}
 