package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(E float32, M float32) float32 {
	// function body
	var N float32 = (E/M)*5 + 1
	return N
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(28.0, 28.0)) // 6.0
	fmt.Println(computeGrade(0.0, 28.0))  // 1.0
}
