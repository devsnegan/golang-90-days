
package main

import "fmt"

func main() {
	// 1. Declare an array
	var numbers [5]int
	fmt.Println("Zero-value array:", numbers)

	// 2. Initialize an array
	scores := [5]int{80, 90, 75, 88, 95}
	fmt.Println("Scores:", scores)

	// 3. Access elements using indexes
	fmt.Println("First score:", scores[0])
	fmt.Println("Last score:", scores[4])

	// 4. Update an element
	scores[2] = 85
	fmt.Println("Updated scores:", scores)

	// 5. Get the array length
	fmt.Println("Number of scores:", len(scores))

	// 6. Iterate using a regular for loop
	fmt.Println("Scores using a for loop:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("Index %d: %d\n", i, scores[i])
	}

	// 7. Iterate using range
	fmt.Println("Scores using range:")
	for index, score := range scores {
		fmt.Printf("Index %d: %d\n", index, score)
	}

	// 8. Calculate the sum
	sum := 0
	for _, score := range scores {
		sum += score
	}
	fmt.Println("Total score:", sum)

	// 9. Calculate the average
	average := float64(sum) / float64(len(scores))
	fmt.Printf("Average score: %.2f\n", average)

	// 10. Arrays are copied when assigned
	first := [3]int{10, 20, 30}
	second := first
	second[0] = 100

	fmt.Println("Original array:", first)
	fmt.Println("Copied array:", second)
}
