package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

type Producter interface {
	GetPrice()
}

type PhysicalProduct struct {
	Name         string
	Price        float64
	Weight       float64
	ShippingCost float64
}

type DigitalProduct struct {
	Name       string
	Price      float64
	FileSize   float64
	LicenseKey string
}

func (p *DigitalProduct) GetPrice() {
	fmt.Println(strconv.FormatFloat(p.Price, 'f', -1, 64))
}

func (p *PhysicalProduct) GetPrice() {
	fmt.Println(strconv.FormatFloat(p.Price, 'f', -1, 64))
}

func (p *PhysicalProduct) ApplyDiscount(discount float64) {
	decrease := p.Price * (discount / 100) // Вычисляем увеличение стоимости
	p.Price = p.Price - decrease
	fmt.Println("Price with discount " + strconv.FormatFloat(p.Price, 'f', -1, 64))

}
func (p *DigitalProduct) ApplyDiscount(discount float64) {
	decrease := p.Price * (discount / 100) // Вычисляем увеличение стоимости
	p.Price = p.Price - decrease
	fmt.Println("Price with discount " + strconv.FormatFloat(p.Price, 'f', -1, 64))
}

func (p *DigitalProduct) GenerateLicenseKey() {
	lenKey := 16
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteKey := make([]byte, lenKey)
	rand.Read(byteKey)

	var keyBuilder strings.Builder
	for i, b := range byteKey {
		if i > 0 && i%4 == 0 {
			// Добавляем дефис каждые 4 символа
			keyBuilder.WriteRune('-')
		}
		keyBuilder.WriteByte(charset[int(b)%len(charset)])
	}
	p.LicenseKey = keyBuilder.String()
	fmt.Println(p.LicenseKey)

}

func (p *PhysicalProduct) CalculateShipping() {
	if p.Price < 100 {
		p.ShippingCost = 15 + p.Weight*0.5
	} else {
		p.ShippingCost = 0
	}
	fmt.Println("Price shipping cost " + strconv.FormatFloat(p.ShippingCost, 'f', -1, 64))
}
func CalculateTotalCost(a []PhysicalProduct, b []DigitalProduct) {
	var totalCost float64
	for i := range a {
		totalCost += a[i].Price + a[i].ShippingCost
	}
	for i := range b {
		totalCost += b[i].Price
	}
	fmt.Println("Total cost is " + strconv.FormatFloat(totalCost, 'f', -1, 64))
}

func main() {
	physBasket := make([]PhysicalProduct, 0)
	digBasket := make([]DigitalProduct, 0)
	physProd1 := PhysicalProduct{"Ticket", 50, 1, 0}
	physProd2 := PhysicalProduct{"Computer", 50, 1, 0}
	physProd3 := PhysicalProduct{"Car", 50, 1, 0}
	physProd4 := PhysicalProduct{"Ganteli", 50, 100, 0}
	digProd := DigitalProduct{"Baldure Gates 3", 100, 240.23, ""}
	physProd1.ApplyDiscount(10)
	physProd2.ApplyDiscount(20)
	physProd3.ApplyDiscount(50)
	physProd4.ApplyDiscount(1)
	physProd1.CalculateShipping()
	physProd2.CalculateShipping()
	physProd3.CalculateShipping()
	physProd4.CalculateShipping()
	physBasket = append(physBasket, physProd1)
	physBasket = append(physBasket, physProd2)
	physBasket = append(physBasket, physProd3)
	physBasket = append(physBasket, physProd4)
	digBasket = append(digBasket, digProd)
	digProd.ApplyDiscount(10)
	digProd.GenerateLicenseKey()
	CalculateTotalCost(physBasket, digBasket)
}
