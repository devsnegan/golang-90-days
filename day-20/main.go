package main

import "fmt"

// 1. Variadic function: accepts zero or more integers.
func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// 2. Normal parameters can come before a variadic parameter.
func greet(message string, names ...string) {
	for _, name := range names {
		fmt.Println(message, name)
	}
}

// 3. Pass a slice using the ... expansion operator.
func calculateAverage(numbers ...int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	total := sum(numbers...)
	return float64(total) / float64(len(numbers))
}

// 4. Accept mixed data types using ...any.
func printDetails(values ...any) {
	for _, value := range values {
		fmt.Println(value)
	}
}

func main() {
	// Multiple arguments
	fmt.Println("Sum:", sum(10, 20, 30))

	// Zero arguments
	fmt.Println("Empty sum:", sum())

	// Passing a slice
	numbers := []int{5, 10, 15, 20}
	fmt.Println("Slice sum:", sum(numbers...))

	// Normal + variadic parameters
	greet("Hello,", "Snegan", "Alex", "John")

	// Average calculation
	fmt.Println("Average:", calculateAverage(10, 20, 30))

	// Mixed data types
	printDetails("Go", 15, true, 99.5)
}
