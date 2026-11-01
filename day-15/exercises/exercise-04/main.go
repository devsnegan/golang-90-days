
package main

import "fmt"

func average(numbers ...int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	total := 0

	for _, number := range numbers {
		total += number
	}

	return float64(total) / float64(len(numbers))
}

func main() {
	result := average(80, 90, 70, 100)

	fmt.Println("Average:", result)
}
