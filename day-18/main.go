package main

import "fmt"

func updateAge(age *int) {
	*age = 30
}

func deposit(balance *float64, amount float64) {
	if balance == nil {
		fmt.Println("Invalid balance pointer")
		return
	}

	if amount <= 0 {
		fmt.Println("Deposit amount must be greater than zero")
		return
	}

	*balance += amount
}

func main() {
	// 1. Normal variable
	age := 25
	fmt.Println("Original age:", age)

	// 2. Get the memory address using &
	agePointer := &age
	fmt.Println("Memory address:", agePointer)

	// 3. Dereference using *
	fmt.Println("Value through pointer:", *agePointer)

	// 4. Modify the original variable through its pointer
	*agePointer = 28
	fmt.Println("Updated age:", age)

	// 5. Pass a pointer to a function
	updateAge(&age)
	fmt.Println("Age after function call:", age)

	// 6. Nil pointer
	var numberPointer *int
	fmt.Println("Nil pointer:", numberPointer)

	if numberPointer == nil {
		fmt.Println("The pointer is nil")
	}

	// 7. Real-world example: bank deposit
	balance := 1000.0
	fmt.Println("\nInitial balance:", balance)

	deposit(&balance, 500)
	fmt.Println("Balance after deposit:", balance)
}
