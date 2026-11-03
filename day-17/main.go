
package main

import (
	"errors"
	"fmt"
)

// Example 1: Named return values with naked return.
func calculate(a, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b

	return
}

// Example 2: Named return values with explicit return.
func calculateSalary(basic, bonus float64) (
	basicSalary float64,
	bonusAmount float64,
	totalSalary float64,
) {
	basicSalary = basic
	bonusAmount = bonus
	totalSalary = basicSalary + bonusAmount

	return basicSalary, bonusAmount, totalSalary
}

// Example 3: Named return value with error.
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

// Example 4: Named return values with zero values.
func getDetails() (name string, age int, active bool) {
	return
}

// Example 5: Product discount calculation.
func calculatePrice(price, discount float64) (
	originalPrice float64,
	discountAmount float64,
	finalPrice float64,
) {
	originalPrice = price
	discountAmount = price * discount / 100
	finalPrice = originalPrice - discountAmount

	return
}

func main() {
	fmt.Println("Named Return Values in Go")
	fmt.Println("-------------------------")

	sum, difference := calculate(20, 10)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)

	basic, bonus, total := calculateSalary(30000, 5000)
	fmt.Println("\nBasic Salary:", basic)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Total Salary:", total)

	balance, err := withdraw(5000, 1500)
	if err != nil {
		fmt.Println("\nTransaction failed:", err)
	} else {
		fmt.Println("\nUpdated Balance:", balance)
	}

	name, age, active := getDetails()
	fmt.Println("\nDefault Name:", name)
	fmt.Println("Default Age:", age)
	fmt.Println("Default Active:", active)

	original, discount, final := calculatePrice(2000, 10)
	fmt.Println("\nOriginal Price:", original)
	fmt.Println("Discount:", discount)
	fmt.Println("Final Price:", final)
}
