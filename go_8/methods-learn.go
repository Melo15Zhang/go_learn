package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

func (x Vertex1) Abs() float64 { // named Abs is alias
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
}

func Abs(v Vertex1) float64 { // named Abs is method
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex1, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	x := Vertex1{3, 4}
	fmt.Println(x.Abs()) //receiver of x
	fmt.Println(Abs(x))  // method

	f := MyFloat(-math.SqrtE)
	fmt.Println(f.Abs()) // receiver of f

	v := Vertex1{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex1{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, *p)
}
