package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func main() {
	c := circle{5}
	fmt.Println(c.area())
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
