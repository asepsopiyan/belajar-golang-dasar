package main

import "fmt"

// Definisikan interface untuk layanan pengiriman
type ShippingService interface {
	CalculateShipping(weight float64, destination string) float64
}

// Implementasikan layanan pengiriman FedEx
type FedEx struct{}

func (f FedEx) CalculateShipping(weight float64, destination string) float64 {
	// Logika penghitungan biaya pengiriman FedEx
	return weight * 2.5 // Misalnya, biaya pengiriman FedEx adalah $2.5 per kg
}

// Implementasikan layanan pengiriman UPS
type UPS struct{}

func (u UPS) CalculateShipping(weight float64, destination string) float64 {
	// Logika penghitungan biaya pengiriman UPS
	return weight * 3.0 // Misalnya, biaya pengiriman UPS adalah $3.0 per kg
}

// Implementasikan layanan pengiriman DHL
type DHL struct{}

func (d DHL) CalculateShipping(weight float64, destination string) float64 {
	// Logika penghitungan biaya pengiriman DHL
	return weight * 2.8 // Misalnya, biaya pengiriman DHL adalah $2.8 per kg
}

// Fungsi untuk mencetak biaya pengiriman berdasarkan layanan pengiriman tertentu
func printShippingCost(service ShippingService, weight float64, destination string) {
	cost := service.CalculateShipping(weight, destination)
	fmt.Printf("Biaya pengiriman dengan %T ke %s adalah $%.2f\n", service, destination, cost)
}

func main() {
	// Buat instance layanan pengiriman
	fedex := FedEx{}
	ups := UPS{}
	dhl := DHL{}

	cost := fedex.CalculateShipping(2.0, "Karawang")
	fmt.Println(cost)

	// Hitung biaya pengiriman untuk setiap layanan
	weight := 5.0
	destination := "New York"
	fmt.Println("Informasi Biaya Pengiriman:")
	printShippingCost(fedex, weight, destination)
	printShippingCost(ups, weight, destination)
	printShippingCost(dhl, weight, destination)
}
