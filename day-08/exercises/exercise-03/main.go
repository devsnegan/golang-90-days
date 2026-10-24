package main

import "fmt"

func main() {
	a, b, c := 25, 40, 15

	if a >= b && a >= c {
		fmt.Println("Largest:", a)
	} else if b >= a && b >= c {
		fmt.Println("Largest:", b)
	} else {
		fmt.Println("Largest:", c)
	}
}