package main

import "fmt"

func main() {
	name := "Joko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Jokowwwwwwi")
	} else {
		fmt.Println("Hallo dek, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
