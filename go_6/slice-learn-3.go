package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dx, dy)
	for i := 0; i < dx; i++ {
		s[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			s[i][j] = uint8((i * j) / 2)
		}
	}
	return s
}

func main() {
	pic := Pic(8, 8)
	for i := 0; i < len(pic); i++ {
		for j := 0; j < len(pic[i]); j++ {
			fmt.Printf("%d\t", pic[i][j])
		}
		fmt.Println()
	}
}
