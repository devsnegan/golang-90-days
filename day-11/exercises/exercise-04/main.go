
package main

import "fmt"

func main() {
	numbers := []int{25, 10, 80, 45, 60, 30}

	largest := numbers[0]

	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}

	fmt.Println("Largest number:", largest)
}
