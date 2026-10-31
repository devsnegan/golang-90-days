package main

import "fmt"

// Function without parameters and return value
func greet() {
	fmt.Println("Welcome to Go Functions")
}

// Function with one parameter
func greetUser(name string) {
	fmt.Println("Hello", name)
}

// Function with multiple parameters
func add(a, b int) {
	fmt.Println("Addition:", a+b)
}

// Function with return value
func multiply(a, b int) int {
	return a * b
}

// Function returning a string
func getName() string {
	return "Snegan"
}

// Function returning a boolean
func isEven(number int) bool {
	return number%2 == 0
}

// Function calling another function
func calculateTotal(price, quantity int) int {
	return price * quantity
}

func printBill(total int) {
	fmt.Println("Total Amount:", total)
}

func main() {
	// 1. Basic function
	greet()

	// 2. Function with parameter
	greetUser("Snegan")

	// 3. Multiple parameters
	add(10, 20)

	// 4. Return value
	result := multiply(5, 4)
	fmt.Println("Multiplication:", result)

	// 5. String return
	fmt.Println("Name:", getName())

	// 6. Boolean return
	fmt.Println("Is Even:", isEven(10))

	// 7. Calling functions from another function
	total := calculateTotal(100, 5)
	printBill(total)
}
