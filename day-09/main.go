
package main

import "fmt"

func main() {
	// 1. Traditional for loop
	fmt.Println("1. Numbers from 1 to 5")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 2. Condition-only loop
	fmt.Println("\n2. Condition-Only Loop")
	count := 1
	for count <= 3 {
		fmt.Println("Count:", count)
		count++
	}

	// 3. range with index and value
	fmt.Println("\n3. Range with Index and Value")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 4. range with values only
	fmt.Println("\n4. Range with Values Only")
	for _, value := range numbers {
		fmt.Println(value)
	}

	// 5. Continue: skip number 3
	fmt.Println("\n5. Continue Statement")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// 6. Break: stop at number 4
	fmt.Println("\n6. Break Statement")
	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	// 7. Sum values using range
	fmt.Println("\n7. Sum of Order Amounts")
	orders := []int{500, 1200, 800, 1500}
	total := 0

	for _, amount := range orders {
		total += amount
	}

	fmt.Println("Number of orders:", len(orders))
	fmt.Println("Total amount:", total)
}
