package main

import "fmt"

func sum(s []int, c chan int) { // take chan in, chan is typed
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // make a channel
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// chan can also take a second arg (buffered chan)
// ch := make(chan int, 100) // buffered channel with capacity 100
// only block sends when buffer is full
// receive blocks when buffer empty
