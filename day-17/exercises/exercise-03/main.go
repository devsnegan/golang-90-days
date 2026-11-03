
package main

import "fmt"

func rectangle(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)

	return
}

func main() {
	area, perimeter := rectangle(10, 5)

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
