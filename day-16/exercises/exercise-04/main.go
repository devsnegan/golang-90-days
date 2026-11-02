
package main

import "fmt"

func checkNumber(number int) (bool, string) {
	if number%2 == 0 {
		return true, "Even number"
	}

	return false, "Odd number"
}

func main() {
	isEven, message := checkNumber(10)

	fmt.Println("Is Even:", isEven)
	fmt.Println("Message:", message)
}
