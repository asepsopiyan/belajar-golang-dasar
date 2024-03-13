package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "EKo"
	middleName = "Kurniawan"
	//lastName = "Khannedy"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
