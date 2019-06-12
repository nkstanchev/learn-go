package main

import (
	"fmt"
)

func ShowcaseSlices() {

	var numbers [5]int
	// how to define array explicitly
	var numbers2 = []float32{1, 2, 3, 4, 21}
	// this is a slice, dynamic array
	var numbersD []int

	numbersD = make([]int, 5, 5)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(numbers, numbers2, numbersD, s)
}
