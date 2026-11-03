
package main

import "fmt"

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
	original, discount, final := calculatePrice(2000, 10)

	fmt.Println("Original Price:", original)
	fmt.Println("Discount:", discount)
	fmt.Println("Final Price:", final)
}
