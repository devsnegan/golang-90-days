package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func Display(value any) {
	fmt.Println("Employee information:", value)
}

func main() {
	employee := Employee{
		Name:   "Snegan",
		Salary: 50000,
	}

	Display(employee)
	Display(101)
	Display("Software Engineer")
}