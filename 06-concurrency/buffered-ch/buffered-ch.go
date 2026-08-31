package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 // two receives in a row
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// channel is "queue-like"
