
package main

import "fmt"

func calculateResult(mark1, mark2, mark3 float64) (
	total float64,
	average float64,
) {
	total = mark1 + mark2 + mark3
	average = total / 3

	return
}

func main() {
	total, average := calculateResult(80, 90, 70)

	fmt.Println("Total Marks:", total)
	fmt.Println("Average Marks:", average)
}
