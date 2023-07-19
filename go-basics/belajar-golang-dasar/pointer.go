package main

import "fmt"

type Address struct {
	City,
	Province,
	Country string
}

// Pointer Function
// before
// func ChangeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

/*
by default variable pass by value
*/

func main() {
	x := 1
	var pX *int = &x

	var y int = *pX

	fmt.Println("pX: ", pX)
	fmt.Println("y: ", y)
}


// func main() {
// 	fmt.Println("# Pointer: ")

// 	// # Pointer
// 	// refer to pointed location

// 	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// 	address2 := address1  // *Address type
// 	address3 := &address2 // <<< & = pointer

// 	address2.City = "Bandung"
// 	address3.City = "Bogor"

// 	fmt.Println("address1: ", address1)
// 	fmt.Println("address2: ", address2)
// 	fmt.Println("address3: ", address3)

// 	// operator *
// 	var address4 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
// 	var address5 *Address = &address4
// 	address5.City = "Bandung"

// 	fmt.Println("address4 - 52: ", address4)
// 	fmt.Println("address5 - 53: ", address5)

// 	/*
// 		# Change the root of pointer to new reference
// 		root 	= adress4,

// 		- Initialy
// 		address5 -> adress4 -> Adress{"Subang"}

// 		- then
// 		address5 -> Address{"Malang"}
// 		address4 -> Address{"Malang"}
// 		address6 -> address4
// 	*/
// 	adress6 := &address4
// 	// change pointer process
// 	*address5 = Address{"Malang", "Jawa Timur", "Indonesia"}
// 	fmt.Println("# after updating pointer: ")

// 	fmt.Println("address4 - 71: ", address4)
// 	fmt.Println("address5 - 72: ", address5)
// 	fmt.Println("adress6: ", adress6)

// 	var address7 *Address = &Address{}
// 	var address8 Address = Address{}

// 	fmt.Println("address7: ", address7)
// 	fmt.Println("address8: ", address8)

// 	// Function New
// 	alamat1 := new(Address)
// 	alamat2 := alamat1

// 	alamat2.Country = "Indonesia"

// 	fmt.Println("alamat1: ", alamat1)
// 	fmt.Println("alamat2: ", alamat2)

// 	// Pointer Function
// 	var alamat = Address{
// 		City:     "Subang",
// 		Province: "Jawa Barat",
// 		Country:  "",
// 	}

// 	// Passing Pointer Argument
// 	// type 1
// 	ChangeCountryToIndonesia(&alamat)

// 	fmt.Println("alamat: ", alamat)

// 	// type 2
// 	var alamaPointer *Address = &alamat
// 	ChangeCountryToIndonesia(alamaPointer)

// 	fmt.Println("alamat: ", alamat)

// }

// // after
// func ChangeCountryToIndonesia(address *Address) {
// 	address.Country = "Indonesia"
// }
