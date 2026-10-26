package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}

	for left, right := 0, len(numbers)-1; left < right; left, right = left+1, right-1 {
		numbers[left], numbers[right] = numbers[right], numbers[left]
	}

	fmt.Println("Reversed array:", numbers)
}
