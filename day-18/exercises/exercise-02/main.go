
package main

import "fmt"

func updateAge(age *int) {
	*age = 30
}

func main() {
	age := 22

	fmt.Println("Before update:", age)

	updateAge(&age)

	fmt.Println("After update:", age)
}
