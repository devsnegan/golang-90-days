package main

import "fmt"

// PrintValue accepts values of any type.
func PrintValue(value any) {
	fmt.Println("Value:", value)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Day 23: Empty Interface and any")

	// Accept different data types.
	PrintValue(100)
	PrintValue("Golang")
	PrintValue(25.5)
	PrintValue(true)

	// Store different values in an any variable.
	var data any

	data = 500
	fmt.Println("Integer:", data)

	data = "Backend Developer"
	fmt.Println("String:", data)

	data = false
	fmt.Println("Boolean:", data)

	// Store a struct.
	student := Student{
		Name: "Snegan",
		Age:  22,
	}

	PrintValue(student)

	// Check the zero value of any.
	var empty any
	fmt.Println("Empty value:", empty)
	fmt.Println("Is nil:", empty == nil)
}
