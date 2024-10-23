package main

import "fmt"

//Задание 1: Модель автомобиля (усложненное)
//Создайте структуру Car с полями Brand, Model, Year, FuelLevel и методом GetInfo(), который возвращает строку с информацией об автомобиле.
//Добавьте методы Drive(distance float64) для уменьшения уровня топлива в зависимости от пройденного расстояния и Refuel(amount float64) для пополнения уровня топлива.
//Также создайте метод NeedsService() для определения необходимости обслуживания автомобиля в зависимости от возраста и пройденного расстояния.

type Car struct {
	Brand        string
	Model        string
	Year         int
	FuelLevel    float64
	FuelLevelMax float64
	Mileage      float64
}

func (car *Car) GetInfo() {
	fmt.Println("Auto -", car.Brand, car.Model, "\nYear of release -", car.Year, "\nCurrent Fuel Level -", car.FuelLevel, "\n")
}

func (car *Car) Drive(distance float64) {
	if distance > car.FuelLevelMax {
		fmt.Println("You cant drive this distance in one try\n")
	} else {
		if distance > car.FuelLevel {
			fmt.Println("You need refuel\n")
		} else {
			car.FuelLevel -= distance
			car.Mileage += distance
		}
	}
}
func (car *Car) Refuel(amount, FuelLevelMax float64) {
	if car.FuelLevelMax > amount+car.FuelLevel {
		car.FuelLevel += amount
	} else {
		car.FuelLevel = car.FuelLevelMax
	}
}
func (car *Car) NeedsService() {
	if car.Mileage > 900 && 2024-car.Year > 20 {
		fmt.Println("You need to pass a technical inspection\n")
	} else {
		if car.Mileage > 10000 {
			fmt.Println("You need to pass a technical inspection\n")
		} else {
			fmt.Println("Ryan Gosling\n")
		}
	}

}

func main() {
	car := Car{
		Brand:        "Opel",
		Model:        "Astra",
		Year:         2000,
		FuelLevel:    500.0,
		FuelLevelMax: 500.0,
		Mileage:      300.0,
	}
	car.NeedsService()
	car.GetInfo()
	car.Drive(500)
	car.GetInfo()
	car.Refuel(500, 500)
	car.GetInfo()
	car.Drive(500)
	car.Drive(1)
	car.GetInfo()
	car.NeedsService()
}
