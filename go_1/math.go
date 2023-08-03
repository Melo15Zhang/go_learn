package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	i := 0

	fmt.Println(" math PI = ", math.Pi)

	fmt.Println(math.Pi)

	for ; i < 10; i++ {
		fmt.Println(" math rand.Int = ", rand.Intn(10))
	}

	fmt.Println(" 36 Sqrt is ", math.Sqrt(36))

}
