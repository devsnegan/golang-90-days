
package main

import "fmt"

func deposit(balance, amount float64) float64 {
	return balance + amount
}

func main() {
	currentBalance := 5000.0
	depositAmount := 1500.0

	updatedBalance := deposit(currentBalance, depositAmount)

	fmt.Println("Current Balance:", currentBalance)
	fmt.Println("Deposit Amount:", depositAmount)
	fmt.Println("Updated Balance:", updatedBalance)
}
