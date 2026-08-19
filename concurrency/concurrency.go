// Package concurrency demonstrates goroutines and channels in Go.
package concurrency

import (
	"sync"
)

// FanOut sends n sequential integers to a channel and returns it.
func FanOut(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

// ParallelSum splits nums into two halves and sums them concurrently.
func ParallelSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mid := len(nums) / 2
	var mu sync.Mutex
	total := 0
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		s := 0
		for _, v := range nums[:mid] {
			s += v
		}
		mu.Lock()
		total += s
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		s := 0
		for _, v := range nums[mid:] {
			s += v
		}
		mu.Lock()
		total += s
		mu.Unlock()
	}()
	wg.Wait()
	return total
}

// Pipeline demonstrates a simple two-stage pipeline.
// Stage 1: square each number. Stage 2: filter even squares.
func Pipeline(nums []int) []int {
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				out <- n * n
			}
		}()
		return out
	}

	filterEven := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				if n%2 == 0 {
					out <- n
				}
			}
		}()
		return out
	}

	in := make(chan int)
	go func() {
		defer close(in)
		for _, n := range nums {
			in <- n
		}
	}()

	squared := square(in)
	filtered := filterEven(squared)

	var result []int
	for n := range filtered {
		result = append(result, n)
	}
	return result
}
