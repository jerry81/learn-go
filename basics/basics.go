// Package basics demonstrates common Go syntax and features.
package basics

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// --- Variables and Constants ---

const Pi = math.Pi

func Variables() {
	// var declaration
	var x int = 10
	var s string = "hello"

	// short declaration
	y := 3.14
	name := "Go"

	// multiple assignment
	a, b := 1, 2
	a, b = b, a

	fmt.Println(x, s, y, name, a, b)
}

// --- Basic Types ---

func BasicTypes() {
	var i int = 42
	var f float64 = 3.14
	var b bool = true
	var s string = "world"
	var r rune = 'A'
	var by byte = 255

	fmt.Printf("int=%d float64=%.2f bool=%t string=%s rune=%c byte=%d\n",
		i, f, b, s, r, by)
}

// --- Control Flow ---

func ControlFlow(n int) string {
	// if / else
	if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	}

	// switch
	switch {
	case n < 10:
		return "small"
	case n < 100:
		return "medium"
	default:
		return "large"
	}
}

// --- Loops ---

func SumTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func CollectEven(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

// --- Functions ---

// Multiple return values
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Variadic function
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Named return values
func MinMax(nums []int) (min, max int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

// --- Closures ---

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// --- Arrays, Slices, Maps ---

func SliceOps() []int {
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	return s[1:4] // [2, 3, 4]
}

func WordCount(sentence string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(sentence) {
		counts[strings.ToLower(word)]++
	}
	return counts
}

// --- Structs ---

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// --- Pointers ---

func Increment(n *int) {
	*n++
}

// --- Error handling ---

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field=%s message=%s", e.Field, e.Message)
}

func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

// --- Defer ---

func WithDefer() string {
	result := []string{}
	defer func() { result = append(result, "deferred") }()
	result = append(result, "first")
	result = append(result, "second")
	_ = result
	return "done"
}
