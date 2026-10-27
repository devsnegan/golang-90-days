
package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}

	numbers = append(numbers, 40)
	numbers = append(numbers, 50, 60)

	fmt.Println("Numbers:", numbers)
	fmt.Println("Length:", len(numbers))
}
