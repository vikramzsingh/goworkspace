// Pointers_Methods_set

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	r := circle{5}
	fmt.Printf("%T\n", r)
	info(&r)

}
