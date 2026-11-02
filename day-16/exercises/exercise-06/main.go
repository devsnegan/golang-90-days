
package main

import (
	"errors"
	"fmt"
)

func withdraw(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("withdrawal amount must be greater than zero")
	}

	if amount > balance {
		return balance, errors.New("insufficient balance")
	}

	return balance - amount, nil
}

func main() {
	balance := 5000.0
	amount := 1500.0

	updatedBalance, err := withdraw(balance, amount)

	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}

	fmt.Println("Withdrawal:", amount)
	fmt.Println("Updated balance:", updatedBalance)
}
