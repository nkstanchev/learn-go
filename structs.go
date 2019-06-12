package main

import "fmt"

type Vertex struct {
	x    int
	y    int
	name string
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func ShowcaseStructs() {
	v := Vertex{1, 2, "example1"} // can't skip values
	var v1 Vertex
	v1.x = 2
	v1.y = 3
	var v2 = Vertex{x: 1, y: 2} // name skipped
	var v3 = new(Vertex)        // v3 is a pointer to an instance of Vertex
	fmt.Println(v, v1, v2, v3)
	// Nested structs

	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	fmt.Println(e.MonthlySalary[0])
	fmt.Println(e.MonthlySalary[1])
	fmt.Println(e.MonthlySalary[2])

}
