package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[2] + fibs[1]
	fibs[4] = fibs[3] + fibs[2]
	// TODO: correct up to index 4 using direct element access

	// TODO: replace 0 with the next Fibonacci number
	fibs = append(fibs, fibs[4]+fibs[3]) // 8
	fibs = append(fibs, fibs[5]+fibs[4]) // 13
	fibs = append(fibs, fibs[6]+fibs[5]) // 21
	// TODO: compute three more Fibonacci numbers and append them
	fibs = append(fibs, fibs[6]+fibs[7])
	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
