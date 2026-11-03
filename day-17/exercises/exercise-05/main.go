
package main

import (
	"errors"
	"fmt"
)

func withdraw(balance, amount float64) (
	updatedBalance float64,
	err error,
) {
	updatedBalance = balance

	if amount <= 0 {
		err = errors.New("amount must be greater than zero")
		return
	}

	if amount > balance {
		err = errors.New("insufficient balance")
		return
	}

	updatedBalance = balance - amount

	return
}

func main() {
	balance, err := withdraw(5000, 1500)

	if err != nil {
		fmt.Println("Transaction failed:", err)
		return
	}

	fmt.Println("Updated Balance:", balance)
}
