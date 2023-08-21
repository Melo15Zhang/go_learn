package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {

	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	s := &Vertex2{3, 4}
	fmt.Println(s.Abs())
	fmt.Println(AbsFunc(*s))
}
