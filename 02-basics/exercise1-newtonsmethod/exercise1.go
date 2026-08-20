package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2.0
	for {
		diff := (z*z - x) / (2 * z)
		if math.Abs(diff) < 1e-9 {
			return z
		}
		z -= diff
	}

}

func main() {
	fmt.Println(Sqrt(2))
}
