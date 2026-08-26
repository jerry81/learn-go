package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v is a negative number", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // this is a conversion expression
	}
	z := x / 2.0
	for {
		diff := (z*z - x) / (2 * z)
		if math.Abs(diff) < 1e-9 {
			return z, nil
		}
		z -= diff
	}
}

func main() {
	if v, err := Sqrt(2); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	if v, err := Sqrt(-2); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
}
