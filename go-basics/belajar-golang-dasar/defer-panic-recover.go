package main

import "fmt"

// # Defer, Panic, Recover

// defer
// keyword to execute function no matter the result
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	/*
	# Place defer at top
	# will be executed, even tho function error
	*/
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result: ", result)
}

// panic
// stop program
func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("Aplikasi berjalan")
}

// recover
// catch panic data, continue func
// run inside defer
func endAppRec() {
	fmt.Println("End App")
	message := recover()

	if message != nil {
		fmt.Println("Terjadi Error ", message)
	}
}
func runAppRec(error bool) {
	defer endAppRec()
	if error {
		panic("App Error")
	}
	fmt.Println("Aplikasi berjalan ")
}

func main() {
	// runApplication(0)
	// runApp(true)
	runAppRec(true)
	fmt.Println("Prove that App still running...")
	fmt.Println("Application Go Rocket")
}
