package main

import (
	"fmt"
	"regexp"
)

// # Package Regexp

func main() {
	fmt.Println("# Package Regexp: ")

	// compile
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	// match
	match1 := regex.MatchString("eko")
	match2 := regex.MatchString("eZo")

	// find all(string, amount)
	search := regex.FindAllString("eko eka ezo", 2) // -1 / all

	fmt.Println("regex: ", regex)
	fmt.Println("match1: ", match1)
	fmt.Println("match2: ", match2)
	fmt.Println("search: ", search)

}
 