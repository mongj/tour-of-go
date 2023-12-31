package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		var z, prev = 1.0, x
		for {
			if math.Abs(z - prev) < 0.000000000000001 {
				return z, nil
			} else {
				prev = z
				z -= (z*z - x) / (2*z)
			}
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
