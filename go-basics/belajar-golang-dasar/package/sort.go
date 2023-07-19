package main

import (
	"fmt"
	"sort"
)

// # Package Sort

type User struct {
	Name 	string
	Age		int
}

type UserSlice []User


func (value UserSlice) Len() int {
	return len(value)
}

// Sort by Age
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	fmt.Println("# Package Sort: ")

	users := []User {
		{"Eko", 30},
		{"Jok", 35},
		{"Bud", 31},
		{"Jon", 25},
	}

	// before
	fmt.Println(users)

	// after
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
 