package main

import (
	"fmt"
	"strings"
)

type Peter interface {
	Move()
	Feed(food string) string
	Speak()
}

func (c *Cat) Move() {
	if c.HungerLevel > 50 {
		c.Speak()
	} else {
		c.HungerLevel += 10
	}
}

func (d *Dog) Move() {
	if d.HungerLevel > 50 {
		d.Speak()
	} else {
		d.HungerLevel += 10
	}
}

func (c Cat) Feed(food string) string {
	if strings.ToLower(food) == "fish" {
		c.HungerLevel = 0
		return fmt.Sprintf("He's full and sleeping")
	} else {
		return fmt.Sprintf("The cat doesn't eat it")
	}
}

func (d Dog) Feed(food string) string {
	if strings.ToLower(food) == "meat" {
		d.HungerLevel = 0
		return fmt.Sprintf("He's full and sleeping")
	} else {
		return fmt.Sprintf("The dog doesn't eat it")
	}
}

func (c Cat) Speak() {
	fmt.Println("Meow")

}

func (d Dog) Speak() {
	fmt.Println("Bark")
}

func (c Cat) Play() {
	if c.Fatigue > 70 {
		fmt.Println("Cat is fall asleep")
	} else {
		c.Fatigue += 25
		fmt.Println("Cap carap")
	}
}

func (d Dog) Play() {
	if d.Fatigue > 70 {
		fmt.Println("Dog is fall asleep")
	} else {
		d.Fatigue += 25
		fmt.Println("He is happy")

	}
}

type Cat struct {
	HungerLevel int
	Fatigue     int
}

type Dog struct {
	HungerLevel int
	Fatigue     int
}

func main() {
	cat := Cat{0, 0}
	dog := Dog{0, 0}
	for i := 0; i < 7; i++ {
		cat.Move()
		dog.Move()
	}
	fmt.Println(dog.Feed("fish"))
	fmt.Println(dog.Feed("meat"))
	fmt.Println(cat.Feed("meat"))
	fmt.Println(cat.Feed("fish"))
	cat.Play()
	dog.Play()
	cat.Fatigue = 100
	dog.Fatigue = 100
	cat.Play()
	dog.Play()
}
