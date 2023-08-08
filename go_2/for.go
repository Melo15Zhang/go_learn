package main

import "fmt"

var sum = 0

func main() {

	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum := 0
	for i := 0; i < 50; i++ {
		sum += i
	}
	fmt.Println(sum)

	for_method()

	for_while()

	for {

	}
}

func for_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func for_method() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
