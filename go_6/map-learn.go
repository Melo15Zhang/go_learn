package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Test center"] = Vertex{
		122.53, 24.39967,
	}
	fmt.Println(m["Test center"])
}
