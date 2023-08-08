package main

import (
	"fmt"
	"math"
)

func main() {
	i := getIfCondition(3)

	if i < 5 {
		fmt.Println("i < 5")
	}

	fmt.Println(sqrt(64), sqrt(-1))

	fmt.Println(pow(2, 3, 10))
	fmt.Println(pow(3, 4, 26))

}

func getIfCondition(i int) (j int) {
	i++
	return i
}

func sqrt(x float64) string {
	if x < 0 {
		s := sqrt(-x)
		if s == "1" {
			return "i"
		} else {
			return fmt.Sprint(s + "i")
		}

	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {

	if v := math.Pow(x, n); v < limit {
		return v
	}
	// return v can not return v
	return limit
}
