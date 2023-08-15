package main

import "fmt"

func main() {

	/** print
	            1
	          1 2
	        1 2 3
	      1 2 3 4
	    1 2 3 4 5
	  1 2 3 4 5 6
	*/
	for i := 1; i < 7; i++ {
		for j := 7 - i; j > 0; j-- {
			fmt.Printf("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
