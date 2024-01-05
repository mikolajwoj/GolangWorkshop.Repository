package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	fmt.Printf("Calculating square number of given %g", x)
	//First use of z
	z -= (z*z - x) / (2 * x)

	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * x)
		fmt.Printf("Attempt number: %v, number received:%g\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
