package main

import "fmt"

func CheckType(value any) {
    switch v := value.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Integer:", v)
    case float64:
        fmt.Println("Float:", v)
    case bool:
        fmt.Println("Boolean:", v)
    default:
        fmt.Printf("Unsupported: %T\n", v)
    }
}

func main() {
    CheckType("Golang")
    CheckType(100)
    CheckType(10.5)
    CheckType(true)
    CheckType([]int{1, 2})
}