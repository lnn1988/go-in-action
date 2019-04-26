package main

import (
	"fmt"
	"math"

)

type geometry interface {
	area() float32
	perimeter() float32
}

type rect struct {
	len, width float32
}

type circle struct {
	radius float32
}

func (r rect)area() float32 {
	return r.len * r.width
}

func (r rect)perimeter() float32  {
	return 2 * (r.len + r.width)
}

func (c circle) area () float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float32 {
	return math.Pi * c.radius * 2
}

func show(name string, params interface{})  {
	switch params.(type) {
	case geometry:
		fmt.Printf("area of %v is %v\n", name, params.(geometry).area())
		fmt.Printf("perimeter of %v is %v\n", name, params.(geometry).perimeter())
	default:
		fmt.Println("Wrong type\n")
	}
}

func main() {
	r := rect{
		len : 1,
		width: 2,
	}
	show("rect", r)

	c := circle{
		radius: 2,
	}
	show("circle", c)
	show("test", "test")
}