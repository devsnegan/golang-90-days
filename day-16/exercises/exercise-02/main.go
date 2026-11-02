
package main

import (
	"errors"
	"fmt"
)

func calculate(a, b float64) (float64, float64, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero is not allowed")
	}

	multiplication := a * b
	division := a / b

	return multiplication, division, nil
}

func main() {
	mul, div, err := calculate(20, 5)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Multiplication:", mul)
	fmt.Println("Division:", div)
}
 