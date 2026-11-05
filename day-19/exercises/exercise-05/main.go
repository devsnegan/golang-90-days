package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func main() {
	books := []Book{
		{
			ID:     1,
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
			Price:  650,
		},
		{
			ID:     2,
			Title:  "Clean Code",
			Author: "Robert C. Martin",
			Price:  750,
		},
		{
			ID:     3,
			Title:  "Designing Data-Intensive Applications",
			Author: "Martin Kleppmann",
			Price:  900,
		},
	}

	fmt.Println("Book Management System")

	for _, book := range books {
		fmt.Println("--------------------")
		fmt.Println("ID:", book.ID)
		fmt.Println("Title:", book.Title)
		fmt.Println("Author:", book.Author)
		fmt.Printf("Price: ₹%.2f\n", book.Price)
	}
}
