
package main

import "fmt"

func main() {
	number := 100

	// Create a pointer to number
	numberPointer := &number

	fmt.Println("Original number:", number)

	// Print memory address
	fmt.Println("Memory address:", numberPointer)

	// Access value through pointer
	fmt.Println("Value using pointer:", *numberPointer)

	// Update value through pointer
	*numberPointer = 200

	fmt.Println("Updated number:", number)
}
