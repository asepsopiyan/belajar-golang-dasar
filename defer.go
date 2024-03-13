package main

import "fmt"

func logging() {
	fmt.Println("Function Selesai Dijalankan")
}

func runApplication() {
	defer logging()
	fmt.Println("RUN APPLICATION..")
}

func main() {
	runApplication()
}
