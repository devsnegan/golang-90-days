package main

import "fmt"

func main() {
	price := 250
	quantity := 4

	total := price * quantity
	discount := 100
	finalAmount := total - discount

	fmt.Println("Total:", total)
	fmt.Println("Discount:", discount)
	fmt.Println("Final Amount:", finalAmount)
	fmt.Println("Remainder of 17 / 5:", 17%5)
} 	