
package main

import "fmt"

func printNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {
	printNumbers(10, 20, 30, 40, 50)
}
