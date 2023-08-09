package main

import "fmt"

func main() {
	defer deferFun()
	fmt.Println("Hello ")

}

func deferFun() {
	v := "World"
	fmt.Println(v)
}
