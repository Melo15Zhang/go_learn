package main

import "fmt"

func main() {

	str := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(str[1:5])
	// slice select array index [x, y] include index-x, exclude index-y
	str1 := str[0:4]
	str2 := str[3:6]
	fmt.Println(str1)
	fmt.Println(str2)
	str2[0] = "g"
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str)

	arr3 := []bool{true, false, true, true}
	fmt.Println(arr3)

	arr4 := []struct {
		i int
		j bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
	}
	fmt.Println(*&arr4)

	arr5 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr5[0:3])
	fmt.Println(arr5[:4])
	fmt.Println(arr5[2:])
	fmt.Println(arr5[:])

	printSlice(arr5[0:3])
	printSlice(arr5[1:])
	printSlice(arr5[:])
}

func printSlice(s []int) {
	fmt.Printf("%v : len=%d cap=%d \n", s, len(s), cap(s))
}
