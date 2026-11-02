
package main

import (
	"errors"
	"fmt"
)

func findMinMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("empty slice")
	}

	min := numbers[0]
	max := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	return min, max, nil
}

func main() {
	numbers := []int{10, 50, 20, 5, 90, 30}

	min, max, err := findMinMax(numbers)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)
}
