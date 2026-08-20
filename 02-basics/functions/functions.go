package main

import "fmt"

func add(x int, y int) int { // type goes after variable name
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
