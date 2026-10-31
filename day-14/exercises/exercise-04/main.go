
package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	number := 15

	if isEven(number) {
		fmt.Println(number, "is Even")
	} else {
		fmt.Println(number, "is Odd")
	}
}
