package main

import "fmt"

func getFullName() (string, string) {
	return "Asep", "Sopiyan"
}

func main() {
	//firstName, _ := getFullName()
	//fmt.Println(firstName)
	//fmt.Println(lastName)

	firstName, _ := getFullName()
	fmt.Println(firstName)
}
