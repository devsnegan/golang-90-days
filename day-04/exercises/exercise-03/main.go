
package main

import "fmt"

func main() {
    const appName = "Payment Service"
    const maxRetries = 3
    const taxRate = 0.18

    fmt.Println("Application:", appName)
    fmt.Println("Maximum Retries:", maxRetries)
    fmt.Println("Tax Rate:", taxRate)
}
