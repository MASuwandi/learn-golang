package main

import (
	"flag"
	"fmt"
)

// # Package Flag

func main() {
	fmt.Println("# Package Flag: ")

	// go run package/flag.go -h

	// (key, default, validation)
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 90, "Put your number")

	// required / parse input
	flag.Parse()

	// return value of address / pointer *
	// without * print result will be pointer / address
	fmt.Println("without *: ", host)
	fmt.Println("*host: ", *host)
	fmt.Println("*username: ", *username)
	fmt.Println("*password: ", *password)
	fmt.Println("*number: ", *number)

	// cli command
	// go run package/flag.go -host=localhost -username=root -password=root
}
 