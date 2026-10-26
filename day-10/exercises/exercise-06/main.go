package main

import "fmt"

func main() {
	original := [4]int{10, 20, 30, 40}
	copied := original

	copied[1] = 200

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)
}
