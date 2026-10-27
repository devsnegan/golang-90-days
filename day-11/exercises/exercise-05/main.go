
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	evenCount := 0
	oddCount := 0

	for _, number := range numbers {
		if number%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Println("Even numbers:", evenCount)
	fmt.Println("Odd numbers:", oddCount)
}
