package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	var l List[int]
	l.val = 1
	item2 := List[int]{val: 2}
	fmt.Println(l.val)
	l.next = &item2
	fmt.Println(l.next.val)
	fmt.Println("should be equal to ", item2.val)
}
