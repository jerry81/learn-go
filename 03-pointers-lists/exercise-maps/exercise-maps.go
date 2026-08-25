package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)

	arr := strings.Fields(s)
	fmt.Println(m)
	for _, word := range arr {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
