package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Marks float64
}

func main() {
	student := Student{
		ID:    101,
		Name:  "Snegan",
		Age:   24,
		Marks: 92.5,
	}

	fmt.Println("Student Details:")
	fmt.Printf("ID: %d\n", student.ID)
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Age: %d\n", student.Age)
	fmt.Printf("Marks: %.2f\n", student.Marks)
}
