package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	gougu := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(gougu(3, 4))
	fmt.Println(compute(gougu))
	fmt.Println(compute(math.Max))

}
