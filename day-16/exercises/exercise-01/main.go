
package main

import "fmt"

func calculate(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, difference := calculate(20, 10)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
