package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7}

	for index, value := range s {
		fmt.Printf("s[%d]=%d\n", index, value)
	}
}
