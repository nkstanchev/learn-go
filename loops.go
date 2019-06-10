package main

import "fmt"

/*
	There is no while/do while loop, just for
*/
func ShowcaseLoops() {

	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	// using range
	for range "Hello" {
		fmt.Println("Hello")
	}
	// while true equivalent
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}
