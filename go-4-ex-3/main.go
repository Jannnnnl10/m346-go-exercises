package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) (float64, float64) {
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return math.NaN(), math.NaN()
	}

	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return x1, x2
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(1, -3, 2)) // 2, 1
	fmt.Println(computeQuadraticFormula(1, 2, 1))  // -1, -1
	fmt.Println(computeQuadraticFormula(1, 1, 1))  // NaN, NaN
}