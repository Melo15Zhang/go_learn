package main

import "fmt"

func main() {

	var arr [4]int
	fmt.Println(arr)

	arr1 := &arr

	arr[0] = 100
	fmt.Println(arr)
	fmt.Println(*arr1)

	var str [2]string
	str[0] = "Hello "
	str[1] = "World!"
	fmt.Println(str[0], str[1])
}
