package main

import "fmt"

type Pricer interface {
	GetPrice() float64
}

type Product struct {
	Name  string
	Price float64
}

func (p Product) GetPrice() float64 {
	return p.Price
}

func CalculateTotalPrice(products []Pricer) float64 {
	total := 0.0
	for _, item := range products {
		total += item.GetPrice()
	}
	return total
}

func main() {
	cart := []Pricer{
		Product{"Iphone13", 2000},
		Product{"IphoneXS", 400},
	}

	totalPrice := CalculateTotalPrice(cart)
	fmt.Println("total belanjaanya dalah Rp.", totalPrice)

}
