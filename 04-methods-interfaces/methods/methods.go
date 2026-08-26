package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // note the syntax - this means Abs() is defined on type Vertex
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// go has no classes!
// abs is a method on the struct Vertex
// vertex is called the receiver
// receiver must be in same package as the method
