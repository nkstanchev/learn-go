package main

import "fmt"

func ShowcaseConditionals() {

	if x := 5; x > 5 {
		fmt.Println("greater than 5")
	} else if x > 3 {
		fmt.Println("greater than 3 less or equal to 5")
	} else {
		fmt.Println("less or equal to 3")
	}

	switch y := 10; {
	case y%2 == 0:
		fmt.Println("even")
		fallthrough
	case y > 10, y > 15:
		fmt.Println("greater than 10")
		fallthrough
	case y > 20:
		fmt.Println("greater than 20")
	case y > 25:
		fmt.Println("greater than 25")
	default:
		fmt.Println("default")

	}
}
