
package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Numbers:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[len(numbers)-1])
}
