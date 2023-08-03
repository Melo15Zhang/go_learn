package main

import "fmt"

func main() {
	x := "StringX"
	y := "StringY"
	fmt.Println(" x, y", x, y)
	x1, y1 := swap(x, y)
	fmt.Println(" x1, y1", x1, y1)

}

func swap(x, y string) (string, string) {
	return y, x
}
