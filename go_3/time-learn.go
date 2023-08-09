package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("now is ", time.Now())

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("today is ", today)
	switch time.Saturday {
	case today + 1:
		fmt.Println("Today.")
	case today + 2:
		fmt.Println("Tomorrow.")
	case today + 3:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	now := time.Now()

	fmt.Println("now day", now.Day())
	fmt.Println("now day", now.Day())
}
