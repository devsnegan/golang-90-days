
package main

import "fmt"

func main() {
	// 1. if statement
	age := 22

	fmt.Println("1. If Statement")
	if age >= 18 {
		fmt.Println("Eligible to register")
	}

	// 2. if-else statement
	balance := 500
	withdrawal := 700

	fmt.Println("\n2. If-Else Statement")
	if withdrawal <= balance {
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Insufficient balance")
	}

	// 3. else-if statement
	score := 82

	fmt.Println("\n3. Else-If Statement")
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else if score >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// 4. Nested if
	isLoggedIn := true
	isAdmin := true

	fmt.Println("\n4. Nested If")
	if isLoggedIn {
		if isAdmin {
			fmt.Println("Admin dashboard available")
		} else {
			fmt.Println("Regular user dashboard available")
		}
	} else {
		fmt.Println("Please log in")
	}

	// 5. Switch with values
	statusCode := 404

	fmt.Println("\n5. Switch Statement")
	switch statusCode {
	case 200:
		fmt.Println("OK")
	case 201:
		fmt.Println("Created")
	case 400:
		fmt.Println("Bad Request")
	case 401:
		fmt.Println("Unauthorized")
	case 404:
		fmt.Println("Not Found")
	default:
		fmt.Println("Other status")
	}

	// 6. Switch with multiple values
	day := "Saturday"

	fmt.Println("\n6. Multiple Values in Switch")
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// 7. Switch without an expression
	temperature := 32

	fmt.Println("\n7. Condition-Based Switch")
	switch {
	case temperature >= 35:
		fmt.Println("Very hot")
	case temperature >= 25:
		fmt.Println("Warm")
	case temperature >= 15:
		fmt.Println("Cool")
	default:
		fmt.Println("Cold")
	}

	// 8. If initializer
	fmt.Println("\n8. If Initializer")
	if limit := 100; balance < limit {
		fmt.Println("Balance is below the limit")
	} else {
		fmt.Println("Balance meets the limit")
	}
}
