package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	employee := Employee{
		ID:         201,
		Name:       "Arun",
		Department: "Backend Development",
		Salary:     45000,
	}

	fmt.Println("Employee Information:")
	fmt.Printf("%+v\n", employee)
}
