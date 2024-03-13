package main

import "fmt"

func main() {
	const firstName = "Asep"
	const lastName = "Sopiyan"

	fmt.Println(firstName)

	//Error
	//firstName = "oke"
	//lastName = "oce"

	const (
		Address = "Tempuran"
		City    = "Karawang"
	)

	fmt.Println(Address)
	fmt.Println(City)
}
