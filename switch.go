package main

import "fmt"

func main() {
	name := "Eza"

	switch name {
	case "EKo":
		fmt.Println("hello fucking eko")
	case "Budi":
		fmt.Println("hello yo budi movfckg")
	default:
		fmt.Println("hello dek, boleh minta wa")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Kepanjangan Satttt")
	case false:
		fmt.Println("Masukkkk")
	}

	name = "Khannedy"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah benar")
	}
}
