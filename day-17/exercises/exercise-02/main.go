
package main

import "fmt"

func calculateSalary(basic, bonus float64) (
	basicSalary float64,
	bonusAmount float64,
	totalSalary float64,
) {
	basicSalary = basic
	bonusAmount = bonus
	totalSalary = basicSalary + bonusAmount

	return
}

func main() {
	basic, bonus, total := calculateSalary(30000, 5000)

	fmt.Println("Basic Salary:", basic)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Total Salary:", total)
}
