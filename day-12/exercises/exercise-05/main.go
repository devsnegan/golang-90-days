package main

import "fmt"

func main() {
	products := map[string]int{
		"Keyboard": 1200,
		"Mouse":    500,
		"Monitor":  8000,
	}

	for product, price := range products {
		fmt.Printf("%s: %d\n", product, price)
	}
}
