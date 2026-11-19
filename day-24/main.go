package main

import "fmt"

func main() {
    var value any = "Golang"

    // One-value type assertion
    result := value.(string)
    fmt.Println("Value:", result)

    // Two-value type assertion
    number, ok := value.(int)

    if ok {
        fmt.Println("Integer:", number)
    } else {
        fmt.Println("Value is not an integer")
    }

    // Type assertion with struct
    type Student struct {
        Name string
        Age  int
    }

    var data any = Student{
        Name: "Snegan",
        Age:  25,
    }

    student, success := data.(Student)

    if success {
        fmt.Println("Student Name:", student.Name)
        fmt.Println("Student Age:", student.Age)
    }
}