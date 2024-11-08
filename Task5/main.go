package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Employee struct {
	Name       string
	Position   string
	Salary     float64
	Experience int
	Post       int
}

type Manager struct {
	employee Employee
	TeamSize int
	Bonus    float64
	Post     int
}

func (m *Manager) CalculateBonus() {
	m.Bonus = float64(m.TeamSize) * 0.25 * float64(m.employee.Experience)
}

func (e *Employee) PromoteEmployee() error {
	if e.Post > 5 {
		return errors.New("Employee has maximum post")
	} else {
		e.Post++
		fmt.Println(e.Name + "has been up to post" + " " + strconv.Itoa(e.Post))
		return nil
	}
}

func (e *Manager) PromoteEmployee() error {
	if e.Post > 5 {
		return errors.New("Employee has maximum post")
	} else {
		e.Post++
		return nil
	}
}

func (e *Manager) CalculateSalary() {
	if e.employee.Experience > 5 {
		e.employee.Salary = (e.employee.Salary + 1.25*float64(e.employee.Experience) + e.Bonus) * 1.1
	} else {
		e.employee.Salary = e.employee.Salary + 1.25*float64(e.employee.Experience) + e.Bonus
	}
}

func (e *Employee) CalculateSalary() {
	if e.Experience > 5 {
		e.Salary = e.Salary + 1.25*float64(e.Experience)*1.1
	} else {
		e.Salary = e.Salary + 1.25*float64(e.Experience)

	}
}

func (e Employee) Vacation() {
	if e.Experience > 3 && e.Post > 3 {
		fmt.Println("An employee" + e.Name + "can rest 30 days a year")
	} else {
		fmt.Println("An employee " + e.Name + "can rest 25 days a year")
	}
}

func (e Employee) GetDetails() {
	fmt.Println("Name:", e.Name, "Position:", e.Position, "Salary:", e.Salary, "Experience:", e.Experience)
}

func main() {
	employee1 := Employee{"Jorik", "Electric", 10000, 2, 0}
	employee1.GetDetails()
	employee1.CalculateSalary()
	employee1.PromoteEmployee()
	employee1.Vacation()
}
