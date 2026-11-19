package main

import "fmt"

type Employee struct {
    Name   string
    Salary float64
}

func main() {
    var value any = Employee{
        Name:   "Arun",
        Salary: 50000,
    }

    employee, ok := value.(Employee)

    if !ok {
        fmt.Println("Not an employee")
        return
    }

    fmt.Println("Name:", employee.Name)
    fmt.Println("Salary:", employee.Salary)
}