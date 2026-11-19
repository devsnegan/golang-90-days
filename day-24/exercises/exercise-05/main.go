package main

import "fmt"

func Process(value any) {
    switch v := value.(type) {
    case string:
        fmt.Println("String:", v)

    case int:
        fmt.Println("Integer:", v)

    case bool:
        fmt.Println("Boolean:", v)

    default:
        fmt.Println("Unsupported type")
    }
}

func main() {
    Process("Golang")
    Process(100)
    Process(true)
    Process(10.5)
}