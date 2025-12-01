package main

import (
	"fmt"
)

// TODO: implement the function computeGrade

func computeGrade(gotPoints, maxPoints float64) float64{
	return 1.0 + (gotPoints/maxPoints)*5.0
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 28.0))
}
