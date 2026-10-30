
package main

import "fmt"

func main() {
	text := "Hello 😊"

	byteCount := len(text)
	runeCount := len([]rune(text))

	fmt.Println("Text:", text)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Rune Count:", runeCount)
}
