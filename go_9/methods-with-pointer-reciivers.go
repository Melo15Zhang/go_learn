package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex3) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	v := &Vertex3{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	v2 := &Vertex3{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs())
	v2.Scale2(5) // not change v2
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs())
}
