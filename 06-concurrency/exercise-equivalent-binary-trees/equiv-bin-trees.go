package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	ch := make(chan int, 10)
	Walk(tree.New(1), ch)
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}
	ch2 := make(chan int, 10)
	func() {
		go Walk(tree.New(2), ch2)
		close(ch2)
	}()

	for j := range ch2 {
		fmt.Println(j)
	}

	fmt.Println("expect true ", Same(tree.New(1), tree.New(1)))
	fmt.Println("expect false ", Same(tree.New(1), tree.New(2)))

}
