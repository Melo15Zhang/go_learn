package main

import (
	"fmt"
	"math"
)

type AbsIntf interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	var absIntf AbsIntf
	var absIntf2 AbsIntf
	f := MyFloat(-math.MaxInt >> 40)
	v := Vertex4{3, 4}

	absIntf = f
	fmt.Println(absIntf.Abs())
	absIntf2 = &v
	fmt.Println(absIntf2.Abs())
	absIntf2 = v // can not call
	fmt.Println(absIntf2.Abs())

}
