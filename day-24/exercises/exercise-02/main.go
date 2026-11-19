package main

import "fmt"

func main() {
    var value any = "Golang"

    number, ok := value.(int)

    fmt.Println("Value:", number)
    fmt.Println("Success:", ok)
}