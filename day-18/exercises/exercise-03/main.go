
package main

import "fmt"

func increaseSalary(salary *float64) {
	*salary = *salary + 5000
}

func main() {
	salary := 25000.0

	fmt.Println("Original salary:", salary)

	increaseSalary(&salary)

	fmt.Println("Updated salary:", salary)
}
