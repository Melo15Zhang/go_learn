package main

import "fmt"

func main() {

	i, j := 42, 2701
	fmt.Println(&i) // 绑定内存地址
	fmt.Println(&j)

	p := &i
	fmt.Println(p)  // 内存地址
	fmt.Println(*p) // 内存地址对应的值
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
