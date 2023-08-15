package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	keys := strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(keys); i++ {
		num := m[keys[i]]
		if num == 0 {
			m[keys[i]] = 1
		} else {
			m[keys[i]] = num + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
