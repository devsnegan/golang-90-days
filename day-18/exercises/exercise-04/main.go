
package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	first := 10
	second := 20

	fmt.Println("Before swapping:")
	fmt.Println("First:", first)
	fmt.Println("Second:", second)

	swap(&first, &second)

	fmt.Println("\nAfter swapping:")
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
}
