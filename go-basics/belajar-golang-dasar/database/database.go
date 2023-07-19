package database

import "fmt"
var connection string

// specific keyword for function name to initialize
func init() {
	fmt.Println("Init Called: ")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}

// # Package Initialization
