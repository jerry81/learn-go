package concurrency

import (
	"testing"
)

func TestFanOut(t *testing.T) {
	// FanOut sends values sequentially, so the order is guaranteed.
	ch := FanOut(5)
	expected := []int{0, 1, 2, 3, 4}
	var got []int
	for v := range ch {
		got = append(got, v)
	}
	if len(got) != len(expected) {
		t.Fatalf("FanOut(5) len = %d, want %d", len(got), len(expected))
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("FanOut[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestParallelSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if got := ParallelSum(nums); got != 55 {
		t.Errorf("ParallelSum = %d, want 55", got)
	}
	if got := ParallelSum(nil); got != 0 {
		t.Errorf("ParallelSum(nil) = %d, want 0", got)
	}
}

func TestPipeline(t *testing.T) {
	// 1,2,3,4,5 -> squares: 1,4,9,16,25 -> even: 4,16
	got := Pipeline([]int{1, 2, 3, 4, 5})
	expected := []int{4, 16}
	if len(got) != len(expected) {
		t.Fatalf("Pipeline len = %d, want %d; got %v", len(got), len(expected), got)
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("Pipeline[%d] = %d, want %d", i, v, expected[i])
		}
	}
}
