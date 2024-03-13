package main

import "fmt"

func main() {
	//var person map[string]string = map[string]string{}
	//person["name"] = "Asep"
	//person["address"] = "Karawang"
	//fmt.Println(person)

	person := map[string]string{
		"name":    "Asep",
		"address": "karawang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Asep"
	book["theme"] = "IT"

	fmt.Println(book)

	delete(book, "theme")

	fmt.Println(book)
}
