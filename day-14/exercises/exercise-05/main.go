
package main

import "fmt"

func calculateArea(length, width float64) float64 {
	return length * width
}

func main() {
	length := 10.0
	width := 5.0

	area := calculateArea(length, width)

	fmt.Println("Length:", length)
	fmt.Println("Width:", width)
	fmt.Println("Area:", area)
}
