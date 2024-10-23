package main

import (
	"fmt"
	"math"
	"slices"
)

type Shaper interface {
	Area() float64
	Perimeter() float64
	IsLargerThan(Shaper) bool
}

type Rectangle struct {
	a float64
	b float64
}

type Circle struct {
	r float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.a + 2*r.b
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) Area() float64 {
	p := (t.a + t.b + t.c) / 2
	fmt.Println(math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c)))
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c Circle) Area() float64 {
	fmt.Println(math.Pi * c.r * c.r)
	return math.Pi * c.r * c.r
}

func (r Rectangle) Area() float64 {
	fmt.Println(r.a * r.b)
	return r.a * r.b
}

func (r Rectangle) IsLargerThan(s Shaper) bool {
	if r.Area() > s.Area() {
		return true
	} else {
		return false
	}
}

func (t Triangle) IsLargerThan(s Shaper) bool {
	if t.Area() > s.Area() {
		return true
	} else {
		return false
	}
}

func (c Circle) IsLargerThan(s Shaper) bool {
	if c.Area() > s.Area() {
		return true
	} else {
		return false
	}
}

func SortShapes(shapes []Shaper) {
	temp := make([]float64, len(shapes))
	result := shapes
	for i, shape := range result {
		temp[i] = shape.Area()
	}
	slices.Sort(temp)
	for i, val := range result {
		if temp[i] <= val.Area() {
			shapes[i] = val
		}
	}
	for _, shape := range shapes {
		fmt.Println(shape)
	}
}

func FilterShapes(shape []Shaper, minArea float64) []Shaper {
	result := make([]Shaper, 0)
	for _, val := range shape {
		if val.Area() > minArea {
			result = append(result, val)
		}
	}
	return result
}

//func (r Rectangle) TransformToSquare() {
//???
//}

func main() {
	r := Rectangle{4, 2}
	t := Triangle{3.0, 4.0, 5.0}
	c := Circle{4}
	r.Area()
	t.Area()
	c.Area()
	shapes := []Shaper{
		Rectangle{4, 2},
		Triangle{3, 4, 5},
		Circle{1},
	}
	SortShapes(shapes)
	fmt.Println(FilterShapes(shapes, 4))
}
