package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7}
	for index, value := range s {
		fmt.Printf("s[%d]=%d\n", index, value)
	}

	m := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range m {
		fmt.Printf("2**[%d]=%d\n", index, value)
	}

	n := make([]int, 10)
	for index := range n {
		n[index] = 1 << uint(index)
	}
	for _, value := range n {
		fmt.Printf("[%d]\n", value)
	}

}
