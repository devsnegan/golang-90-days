package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Mango", "Orange"}

	fmt.Println("Fruits:", fruits)
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Last fruit:", fruits[len(fruits)-1])
	fmt.Println("Total fruits:", len(fruits))
}
