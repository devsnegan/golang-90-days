package main

import "fmt"

func Log(value any) {
    switch v := value.(type) {
    case nil:
        fmt.Println("Value is nil")

    case string:
        fmt.Println("Message:", v)

    case int:
        fmt.Println("Integer:", v)

    case bool:
        fmt.Println("Boolean:", v)

    case float64:
        fmt.Println("Amount:", v)

    default:
        fmt.Printf("Unsupported type: %T\n", v)
    }
}

func main() {
    Log("User logged in")
    Log(101)
    Log(true)
    Log(2500.50)
    Log(nil)
    Log([]int{1, 2, 3})
}