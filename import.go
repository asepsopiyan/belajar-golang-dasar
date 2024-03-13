package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Asep")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // tidak bisa akses
	//fmt.Println(helper.sayGoodbye("Asep")) // tidak bisa akses
}
