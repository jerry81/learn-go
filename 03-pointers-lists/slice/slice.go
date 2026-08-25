package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // notice no size
	fmt.Printf("len is %d", len(s))
	fmt.Println(s)
}

// arrays actually rare in go code.
// they are like freeze in ruby
// "zero" value of slice is nil
