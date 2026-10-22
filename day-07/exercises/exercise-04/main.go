package main

import "fmt"

func main() {
	balance := 1000

	balance += 500
	fmt.Println("After deposit:", balance)

	balance -= 200
	fmt.Println("After withdrawal:", balance)

	balance++
	fmt.Println("After increment:", balance)

	balance--
	fmt.Println("After decrement:", balance)
}