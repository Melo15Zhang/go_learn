package main

import "fmt"

func main() {

	vr := Vertex{1, 2}
	fmt.Println(vr.X)
	fmt.Println(vr.Y)
	vr.X = 5
	fmt.Println(vr.X)

	p := &vr
	p.X = 1e1
	p.Y = 3222

	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(v1, v2, v3, *v4)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 4}
	v3 = Vertex{}
	v4 = &v1
)

type Vertex struct {
	X int
	Y int
}
