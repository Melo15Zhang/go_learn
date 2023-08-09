package main

import (
	"fmt"
	"math/rand"
)

func main() {

	great := 0
	good := 0
	notBad := 0
	bad := 0
	for i := 0; i < 100; i++ {
		score := getScore()
		switch {
		case score >= 90:
			fmt.Printf("The score %d is Great! level A+\n", score)
			great++
		case score >= 80 && score < 90:
			fmt.Printf("The score %d Good! level A\n", score)
			good++
		case score >= 70 && score < 80:
			fmt.Printf("The score %d not Bad! level B\n", score)
			notBad++
		case score < 70:
			fmt.Printf("The score %d Bad! level C\n", score)
			bad++
		}
	}
	fmt.Printf("The score A+ appears %d times \n", great)
	fmt.Printf("The score A appears %d times \n", good)
	fmt.Printf("The score B appears %d times \n", notBad)
	fmt.Printf("The score C appears %d times \n", bad)
}

func getScore() (score int) {
	return rand.Intn(100)
}
