
package main

import "fmt"

func main() {
	text := "Goக😊"

	runes := []rune(text)

	fmt.Println("Original:", text)
	fmt.Println("Rune Values:", runes)
	fmt.Println("Rune Count:", len(runes))

	for index, value := range runes {
		fmt.Printf("Index: %d, Character: %c\n", index, value)
	}
}
