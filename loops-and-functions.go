package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z, prev = 1.0, x
	for {
		fmt.Println("z:", z)
		if math.Abs(z - prev) < 0.000000000000001 {
			return z
		} else {
			prev = z
			z -= (z*z - x) / (2*z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
