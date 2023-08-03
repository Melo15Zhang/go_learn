package main

import "fmt"

func main() {

	// 3 + 7 = ?
	k := 3
	n := 7
	fmt.Println(" 3 + 7 is", add(k, n))
}

func add(k, n int) int {
	return k + n
}
