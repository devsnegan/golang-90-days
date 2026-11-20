package main

import "fmt"

type Employee struct {
    Name   string
    Salary float64
}

type Student struct {
    Name string
    Age  int
}

func Display(value any) {
    switch v := value.(type) {
    case Employee:
        fmt.Println("Employee:", v.Name, v.Salary)
    case Student:
        fmt.Println("Student:", v.Name, v.Age)
    default:
        fmt.Println("Unknown record")
    }
}

func main() {
    Display(Employee{Name: "Arun", Salary: 50000})
    Display(Student{Name: "Kumar", Age: 22})
}