package main

import "fmt"

var c, python, java = true, false, false

func main() {
	var i, j = 9, 10
	c = true
	fmt.Println(i, j, c, python, java)

	k := "sass"
	fmt.Println(k)
}
