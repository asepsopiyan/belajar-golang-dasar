package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var asep Customer

	asep.Name = "Asep Sopiyan"
	asep.Address = "Indonesia"
	asep.Age = 30

	fmt.Println(asep)
	fmt.Println(asep.Name)
	fmt.Println(asep.Address)
	fmt.Println(asep.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Malaysia",
		Age:     23,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "America", 44}
	fmt.Println(budi)

	budi.sayHello("Agus")
	asep.sayHello("Agus")
	joko.sayHello("Agus")
}
