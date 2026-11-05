package main

import "fmt"

// Define a Student struct
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// Define a Product struct
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Define an Address struct
type Address struct {
	City    string
	State   string
	Country string
}

// Define an Employee struct
type Employee struct {
	ID      int
	Name    string
	Salary  float64
	Address Address
}

func main() {
	// 1. Create a struct using keyed fields
	student := Student{
		ID:    101,
		Name:  "Snegan",
		Age:   24,
		Grade: "A",
	}

	fmt.Println("Student Details:")
	fmt.Println("ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Grade:", student.Grade)

	// 2. Update a struct field
	student.Age = 25
	fmt.Println("\nUpdated Age:", student.Age)

	// 3. Create a struct using positional values
	product := Product{1, "Laptop", 55000, 10}

	fmt.Println("\nProduct Details:")
	fmt.Printf("%+v\n", product)

	// 4. Struct zero values
	var newStudent Student
	fmt.Println("\nNew Student (zero values):")
	fmt.Printf("%+v\n", newStudent)

	// 5. Nested structs
	employee := Employee{
		ID:     501,
		Name:   "Arun",
		Salary: 45000,
		Address: Address{
			City:    "Chennai",
			State:   "Tamil Nadu",
			Country: "India",
		},
	}

	fmt.Println("\nEmployee Details:")
	fmt.Printf("%+v\n", employee)
	fmt.Println("Employee City:", employee.Address.City)
}
