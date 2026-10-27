
package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Println("Sum:", sum)
	fmt.Printf("Average: %.2f\n", average)
}
