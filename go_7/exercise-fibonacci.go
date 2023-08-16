package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0, f1, value := 1, 1, 0
	index := 0
	return func() int {
		index++
		if index == 1 {
			return 0
		} else if index == 2 || index == 3 {
			return 1
		} else {
			value, f0 = f0+f1, f1
			f1 = value
			return value
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,", f())
	}
}
