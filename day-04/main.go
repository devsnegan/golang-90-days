
package main

import "fmt"

func main() {
    var name string = "Snegan"
    var role = "Software Developer"

    age := 25
    isLearning := true

    const language = "Golang"
    const maxRetries = 3

    fmt.Println("===== USER PROFILE =====")
    fmt.Println("Name:", name)
    fmt.Println("Role:", role)
    fmt.Println("Age:", age)
    fmt.Println("Learning:", language)
    fmt.Println("Is Learning:", isLearning)
    fmt.Println("Maximum Retries:", maxRetries)

    age = 26
    fmt.Println("Updated Age:", age)
}
