package main

import (
	"fmt"

	"github.com/jerry81/learn-go/basics"
)

func main() {
	fmt.Println("Add(4, 6):", basics.Add(4, 6))
	fmt.Println("IsEven(11):", basics.IsEven(11))
	fmt.Println("WordFrequency:", basics.WordFrequency([]string{"go", "go", "syntax"}))

	result, err := basics.Divide(20, 4)
	if err != nil {
		fmt.Println("Divide error:", err)
		return
	}
	fmt.Println("Divide(20, 4):", result)
}
