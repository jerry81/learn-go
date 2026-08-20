package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// basically, defer evaluates immediately but executed after the return
// application: cleanup - kind of like finally?
