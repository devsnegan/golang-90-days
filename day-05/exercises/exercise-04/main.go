package main

import "fmt"

func main() {
	text := "Go அ"

	fmt.Println("Text:", text)
	fmt.Println("Byte length:", len(text))
	fmt.Println("Rune count:", len([]rune(text)))

	fmt.Println("Characters:")
	for _, r := range text {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}