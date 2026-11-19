package main

import "fmt"

func main() {
    var value any = 100

    number := value.(int)

    fmt.Println("Integer:", number)
}