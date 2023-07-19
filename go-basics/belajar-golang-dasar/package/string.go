package main

import (
	"fmt"
	"strings"
)

// # Package String

func main() {
	fmt.Println("# Package String: ")

	// Trim
	trim := strings.Trim("  Eko Khanedy  ", " ")

	// ToLower
	toLower := strings.ToLower("Eko Khanedy")

	// ToUpper
	toUpper := strings.ToUpper("Eko Khanedy")

	// Split(string, separator) return array
	split := strings.Split("Eko Khanedy", " ")

	// Contains(string, search)
	containsEko := strings.Contains("Eko Khanedy", "Eko")
	containsBudi := strings.Contains("Eko Khanedy", "Budi")

	// ReplaceAll(string, search, replace)
	replaceAll := strings.ReplaceAll("Eko Khanedy", "Eko", "Budi")

	fmt.Println("trim: ", trim)
	fmt.Println("toLower: ", toLower)
	fmt.Println("toUpper: ", toUpper)
	fmt.Println("split: ", split)
	fmt.Println("containsEko: ", containsEko)
	fmt.Println("containsBudi: ", containsBudi)
	fmt.Println("replaceAll: ", replaceAll)
}
 