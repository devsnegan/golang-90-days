
package main

import (
	"errors"
	"fmt"
)

// Example 1: Return two values.
func calculate(a, b int) (int, int) {
	sum := a + b
	difference := a - b

	return sum, difference
}

// Example 2: Return a result and an error.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return a / b, nil
}

// Example 3: Return minimum and maximum.
func findMinMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("numbers slice cannot be empty")
	}

	min := numbers[0]
	max := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
	}

	return min, max, nil
}

// Example 4: Ignore an unwanted return value.
func getUserDetails() (string, int) {
	return "Alex", 25
}

// Example 5: Return username and error.
func getUser(id int) (string, error) {
	users := map[int]string{
		101: "Alex",
		102: "John",
		103: "David",
	}

	name, exists := users[id]

	if !exists {
		return "", errors.New("user not found")
	}

	return name, nil
}

func main() {
	fmt.Println("Multiple Return Values in Go")
	fmt.Println("-----------------------------")

	sum, difference := calculate(20, 10)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)

	result, err := divide(20, 5)
	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	numbers := []int{10, 50, 20, 5, 90, 30}

	min, max, err := findMinMax(numbers)
	if err != nil {
		fmt.Println("MinMax error:", err)
	} else {
		fmt.Println("Minimum:", min)
		fmt.Println("Maximum:", max)
	}

	// Ignore the age and receive only the username.
	username, _ := getUserDetails()
	fmt.Println("Username:", username)

	name, err := getUser(102)
	if err != nil {
		fmt.Println("User error:", err)
	} else {
		fmt.Println("Found user:", name)
	}
}
