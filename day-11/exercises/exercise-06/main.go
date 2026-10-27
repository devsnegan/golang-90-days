
package main

import "fmt"

func main() {
	original := []int{10, 20, 30, 40}

	copied := make([]int, len(original))
	copy(copied, original)

	copied[0] = 100
	copied[1] = 200

	fmt.Println("Original slice:", original)
	fmt.Println("Copied slice:", copied)
}
