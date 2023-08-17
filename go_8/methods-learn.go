package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (x Vertex) Abs() float64 { // named Abs is alias
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
}

func Abs(v Vertex) float64 { // named Abs is method
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	x := Vertex{3, 4}
	fmt.Println(x.Abs()) //receiver of x
	fmt.Println(Abs(x))  // method

	f := MyFloat(-math.SqrtE)
	fmt.Println(f.Abs()) // receiver of f
}
