package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	// function body
	var hypotenuse = math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return float64(hypotenuse)
}

func main() {
	// TODO: call the function computeHypotenuse

	fmt.Print(computeHypotenuse(3, 4)) // 5
}
