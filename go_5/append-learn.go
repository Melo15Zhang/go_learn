package main

import "fmt"

func main() {
	stringSliceAppend()
	intSliceAppend()
}

func stringSliceAppend() {
	var s []string
	printSliceString(s)

	s = append(s, "A")
	printSliceString(s)

	s = append(s, "B")
	printSliceString(s)

	s = append(s, "C", "D", "E")
	printSliceString(s)
}

func intSliceAppend() {
	var i []int
	printSliceInt(i)

	i = append(i, 0)
	printSliceInt(i)

	i = append(i, 1)
	printSliceInt(i)

	i = append(i, 2)
	printSliceInt(i)

	i = append(i, 3)
	printSliceInt(i)

	i = append(i, 5)
	printSliceInt(i)
}

func printSliceString(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceInt(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
