package main

import (
	"fmt"
	"os"
)

// # Package os

func main() {
	fmt.Println("# Package os: ")

	// Args
	args := os.Args
	
	fmt.Println("Args: ", args)
	fmt.Println("Args[0]: ", args[0])

	// go run package/package-os.go [arg/param]
	// fmt.Println("Args: ", args[1])

	// run hostname in shell
	// hostname in code

	hostname, err := os.Hostname()

	if(err == nil) {
		fmt.Println("hostname: ", hostname)
	} else {
		fmt.Println("err.Error(): ", err.Error())
	}

	// get env
	// export shell command
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println("username: ", username)
	fmt.Println("password: ", password)

}
 