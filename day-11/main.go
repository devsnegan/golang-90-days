package main

import (
	"fmt"
)

func main() {
	// 1. Create a slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", numbers)

	// 2. Access elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[len(numbers)-1])

	// 3. Update an element
	numbers[1] = 200
	fmt.Println("Updated slice:", numbers)

	// 4. Find length and capacity
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// 5. Append elements
	numbers = append(numbers, 60, 70)
	fmt.Println("After append:", numbers)

	// 6. Iterate using range
	fmt.Println("Slice elements:")
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// 7. Calculate sum and average
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	average := float64(sum) / float64(len(numbers))
	fmt.Println("Sum:", sum)
	fmt.Printf("Average: %.2f\n", average)

	// 8. Create a slice using make
	scores := make([]int, 3, 5)
	scores[0] = 80
	scores[1] = 90
	scores[2] = 85
	fmt.Println("Scores:", scores)
	fmt.Println("Scores length:", len(scores))
	fmt.Println("Scores capacity:", cap(scores))

	// 9. Create a subslice
	values := []int{10, 20, 30, 40, 50}
	part := values[1:4]
	fmt.Println("Original values:", values)
	fmt.Println("Subslice:", part)

	// 10. Copy a slice
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 100

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// 11. Nil slice versus empty slice
	var nilSlice []int
	emptySlice := []int{}

	fmt.Println("Nil slice:", nilSlice)
	fmt.Println("Nil slice is nil:", nilSlice == nil)
	fmt.Println("Empty slice:", emptySlice)
	fmt.Println("Empty slice is nil:", emptySlice == nil)

	// 12. Append to a nil slice
	var newNumbers []int
	newNumbers = append(newNumbers, 10, 20, 30)
	fmt.Println("Slice after appending:", newNumbers)
}
