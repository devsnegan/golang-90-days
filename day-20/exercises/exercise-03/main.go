
package main

import "fmt"

func findLargest(numbers ...int) (int, bool) {
	if len(numbers) == 0 {
		return 0, false
	}

	largest := numbers[0]

	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}

	return largest, true
}

func main() {
	result, ok := findLargest(10, 50, 30, 90, 20)

	if !ok {
		fmt.Println("No numbers provided")
		return
	}

	fmt.Println("Largest:", result)
}
