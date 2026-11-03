
package main

import "fmt"

func calculate(a, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b

	return
}

func main() {
	sum, difference := calculate(20, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
