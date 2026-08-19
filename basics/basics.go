package basics

import "fmt"

func Add(a, b int) int {
	return a + b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func WordFrequency(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
