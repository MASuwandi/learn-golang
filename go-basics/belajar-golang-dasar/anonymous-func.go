package main

import "fmt"

// # Anonymous Function

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked: ", name)
	} else {
		fmt.Println("Welcome: ", name)
	}
}



func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("go", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}

// Manual way

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }
