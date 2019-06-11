package main

import (
	"fmt"
)

func main() {

	ShowcaseOperators()
	ShowcaseVariables()

	// import statements are preprocessors
	// the main package is the starting point of the program
	// IMPORTANT: A name is exported if it starts with capital letter
	// array definition
	var numbers [5]int
	// how to define array explicitly
	var numbers2 = []float32{1, 2, 3, 4, 21}
	// this is a slice, dynamic array
	var numbersD []int

	numbersD = make([]int, 5, 5)
	fmt.Println(numbers)
	fmt.Println(numbers2)
	// slices23
	fmt.Println(numbersD)
	ShowcaseChannels()
	// fmt.Println(slices23())
	// structures
	// v := Vertex{1, 2}
	// v.x = 3
	// fmt.Println(v)
}
