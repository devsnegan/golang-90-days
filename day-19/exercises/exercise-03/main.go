package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func main() {
	product := Product{
		ID:    301,
		Name:  "Keyboard",
		Price: 1500,
		Stock: 20,
	}

	fmt.Println("Before Update:")
	fmt.Printf("%+v\n", product)

	product.Price = 1800
	product.Stock = 15

	fmt.Println("\nAfter Update:")
	fmt.Printf("%+v\n", product)
}
