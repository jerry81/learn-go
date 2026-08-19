package basics

import (
	"math"
	"testing"
)

func TestControlFlow(t *testing.T) {
	cases := []struct {
		input    int
		expected string
	}{
		{-1, "negative"},
		{0, "zero"},
		{5, "small"},
		{50, "medium"},
		{200, "large"},
	}
	for _, tc := range cases {
		got := ControlFlow(tc.input)
		if got != tc.expected {
			t.Errorf("ControlFlow(%d) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}

func TestSumTo(t *testing.T) {
	if got := SumTo(10); got != 55 {
		t.Errorf("SumTo(10) = %d, want 55", got)
	}
	if got := SumTo(0); got != 0 {
		t.Errorf("SumTo(0) = %d, want 0", got)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil || result != 5.0 {
		t.Errorf("Divide(10, 2) = %v, %v; want 5.0, nil", result, err)
	}
	_, err = Divide(1, 0)
	if err == nil {
		t.Error("Divide(1, 0) expected error, got nil")
	}
}

func TestSum(t *testing.T) {
	if got := Sum(1, 2, 3, 4); got != 10 {
		t.Errorf("Sum(1,2,3,4) = %d, want 10", got)
	}
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
}

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if min != 1 || max != 9 {
		t.Errorf("MinMax = %d, %d; want 1, 9", min, max)
	}
	min, max = MinMax(nil)
	if min != 0 || max != 0 {
		t.Errorf("MinMax(nil) = %d, %d; want 0, 0", min, max)
	}
}

func TestCounter(t *testing.T) {
	c := Counter()
	for i := 1; i <= 3; i++ {
		if got := c(); got != i {
			t.Errorf("counter call %d = %d, want %d", i, got, i)
		}
	}
}

func TestWordCount(t *testing.T) {
	counts := WordCount("go is go")
	if counts["go"] != 2 {
		t.Errorf("WordCount: 'go' = %d, want 2", counts["go"])
	}
	if counts["is"] != 1 {
		t.Errorf("WordCount: 'is' = %d, want 1", counts["is"])
	}
}

func TestShapes(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if r.Area() != 12 {
		t.Errorf("Rectangle area = %v, want 12", r.Area())
	}
	if r.Perimeter() != 14 {
		t.Errorf("Rectangle perimeter = %v, want 14", r.Perimeter())
	}

	c := Circle{Radius: 1}
	if got := c.Area(); math.Abs(got-math.Pi) > 1e-9 {
		t.Errorf("Circle area = %v, want Pi", got)
	}
}

func TestIncrement(t *testing.T) {
	n := 5
	Increment(&n)
	if n != 6 {
		t.Errorf("Increment: got %d, want 6", n)
	}
}

func TestValidateAge(t *testing.T) {
	if err := ValidateAge(25); err != nil {
		t.Errorf("ValidateAge(25) unexpected error: %v", err)
	}
	if err := ValidateAge(-1); err == nil {
		t.Error("ValidateAge(-1) expected error")
	}
	if err := ValidateAge(150); err != nil {
		t.Errorf("ValidateAge(150) unexpected error: %v", err)
	}
	if err := ValidateAge(151); err == nil {
		t.Error("ValidateAge(151) expected error")
	}
}
