package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	fmt.Printf("Calculating square number of given %g\n", x)

	//First use of z
	//z -= (z*z - x) / (2 * x)

	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Attempt number: %v, number received:%g\n", i, z)
		if math.Abs(math.Sqrt(x)-z) < 1e-9 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
