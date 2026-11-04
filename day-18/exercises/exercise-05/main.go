
package main

import "fmt"

func deposit(balance *float64, amount float64) {
	// Check whether pointer is nil
	if balance == nil {
		fmt.Println("Error: Invalid balance pointer")
		return
	}

	// Validate deposit amount
	if amount <= 0 {
		fmt.Println("Error: Deposit amount must be greater than zero")
		return
	}

	// Update original balance
	*balance += amount

	fmt.Println("Deposit successful")
}

func main() {
	balance := 10000.0

	fmt.Println("Initial balance:", balance)

	deposit(&balance, 2500)

	fmt.Println("Final balance:", balance)

	// Test invalid deposit
	deposit(&balance, -500)

	// Test nil pointer
	deposit(nil, 1000)
}
