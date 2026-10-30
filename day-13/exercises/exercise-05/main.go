
package main

import "fmt"

func main() {
	text := "Hello Go"

	for index, character := range text {
		fmt.Printf("Index: %d, Character: %c\n", index, character)
	}
}
