package main

import (
	"errors"
	"fmt"
	"sort"
)

type Animal struct {
	Name    string
	Spices  string
	Age     int
	Fatigue int
}

type Zoo struct {
	animals []Animal
}

func (z *Zoo) AddAnimal(a Animal) {
	z.animals = append(z.animals, a)
}

func (z *Zoo) RemoveAnimal(a Animal) string {
	for i := range z.animals {
		if z.animals[i] == a {
			z.animals = append(z.animals[:i], z.animals[i+1:]...)
			return "You killed " + a.Name
		}
	}
	return "All " + a.Name + " have been killed in your zoo"
}

func (a Animal) Speak() {
	fmt.Println("skibiditoilet")
}

func (z Zoo) FindOldestAnimal() {
	sort.Slice(z.animals, func(i, j int) bool { return z.animals[i].Age < z.animals[j].Age })
	fmt.Println(z.animals[len(z.animals)-1])
}

func (z *Zoo) FeedAllAnimals(food string) error {
	for i := range z.animals {
		if z.animals[i].Spices == "mamal" && (food == "meat" || food == "tomatoes") {
			fmt.Println(z.animals[i])
		} else {
			return errors.New(z.animals[i].Name + " cant eat " + food)
		}
	}
	return nil
}

func (z Zoo) CountBySpices(species string) int {
	var counter int
	for i := range z.animals {
		if z.animals[i].Spices == species {
			counter++
		}
	}
	return counter
}

func (z Zoo) ScheduleEvent(name string, animals []Animal) []Animal {
	members := make([]Animal, 0)
	for i := range animals {
		if animals[i].Fatigue > 70 {
			fmt.Println(animals[i].Name + " is very low hp,he will didnt participate event")
		} else if animals[i].Age > 15 {
			fmt.Println(animals[i].Name + " is very old,he will didnt participate event")
		} else if animals[i].Spices == "Delphine" {
			fmt.Println(animals[i].Name + " live in water,he will didnt participate event")
		} else {
			members = append(members, animals[i])
		}
	}
	return members
}

func main() {
	z := Zoo{}
	animal := Animal{Name: "Tiger", Spices: "mamal", Age: 5, Fatigue: 30}
	z.AddAnimal(animal)
	z.AddAnimal(Animal{Name: "Elephant", Spices: "mamal", Age: 10, Fatigue: 50})
	z.AddAnimal(Animal{Name: "Tiger", Spices: "mamal", Age: 5, Fatigue: 30})
	z.AddAnimal(Animal{Name: "Parrot", Spices: "Delphine", Age: 2, Fatigue: 0})
	err := z.FeedAllAnimals("meat")
	fmt.Println(err)
	fmt.Println(z.CountBySpices("bird"))
	fmt.Println(z.ScheduleEvent("Wild Event", z.animals))
}
