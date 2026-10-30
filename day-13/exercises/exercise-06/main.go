
package main

import "fmt"

func main() {
	text := "Golang"

	runes := []rune(text)

	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	reversed := string(runes)

	fmt.Println("Original:", text)
	fmt.Println("Reversed:", reversed)
}
