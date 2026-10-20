package main

import "fmt"

func main() {
	var age int = 25
	var price float64 = 99.5

	result := float64(age) + price

	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println("Result:", result)
}
