package main

//example for https://golang.google.cn/tour/moretypes/18 Exercise
import "golang.org/x/tour/pic"

func Pic1(dx, dy int) [][]uint8 {
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
	pic.Show(Pic1)
}
