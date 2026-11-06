
package main

import "fmt"

func calculateSum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := calculateSum(10, 20, 30, 40)

	fmt.Println("Total:", result)
}
