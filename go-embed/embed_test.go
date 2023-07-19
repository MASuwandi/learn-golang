package main_test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// Embed String
//
//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println("version: ", version)
}

// Embed []byte
//
//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	// write file png
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("logo: ", logo)
}

// Embed Multiple Files

// Patterns may not contain ‘.’ or ‘..’ or empty path elements, nor may they begin or end with a slash.
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	// need read file to read file
	// simple one by one
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println("a: ", string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println("b: ", string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println("c: ", string(c))

	fmt.Println("TestMultipleFiles Done: ")
}

// Embed Path Matcher

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	// dynamic
	dirEntries, _ := path.ReadDir("files")
	fmt.Println("dirEntries: ", dirEntries)

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("file name: ", entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("file: ", string(file))
		}
	}

	fmt.Println("TestPathMatcher Done: ")
}
