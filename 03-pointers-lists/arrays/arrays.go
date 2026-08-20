package main

import "fmt"

func main() {
	var a [2]string // declare size up front
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a) // println handles arrays

	primes := [6]int{2, 3, 5, 7, 11, 13} // inline initializiation
	fmt.Println(primes)
}

// arrays can't be resized
