
package main

import "fmt"

func calculateTotal(prices ...float64) float64 {
	total := 0.0

	for _, price := range prices {
		total += price
	}

	return total
}

func main() {
	total := calculateTotal(100.50, 250.00, 149.50)

	fmt.Printf("Shopping Cart Total: %.2f\n", total)
}
