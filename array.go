package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Asep"
	names[1] = "Sopiyan"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var freeValues = [...]int{
		20,
		90,
		88,
		76,
	}

	fmt.Println(len(freeValues))

}
