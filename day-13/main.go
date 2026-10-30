package main

import "fmt"

func main() {
	// 1. Basic string
	text := "Goக😊"

	fmt.Println("Original String:", text)

	// 2. String length in bytes
	fmt.Println("\nString length in bytes:", len(text))

	// 3. Convert string to byte slice
	bytes := []byte(text)

	fmt.Println("\nByte Slice:", bytes)
	fmt.Println("Byte Count:", len(bytes))

	// 4. Convert string to rune slice
	runes := []rune(text)

	fmt.Println("\nRune Slice:", runes)
	fmt.Println("Rune Count:", len(runes))

	// 5. Access individual bytes
	fmt.Println("\nFirst Byte:", bytes[0])

	// 6. Access individual runes
	fmt.Printf("First Rune: %c\n", runes[0])
	fmt.Printf("Tamil Rune: %c\n", runes[2])
	fmt.Printf("Emoji Rune: %c\n", runes[3])

	// 7. Iterate using range
	fmt.Println("\nIterating over String:")

	for index, character := range text {
		fmt.Printf("Byte Index: %d, Character: %c\n", index, character)
	}

	// 8. Iterate over byte slice
	fmt.Println("\nIterating over Bytes:")

	for index, value := range bytes {
		fmt.Printf("Index: %d, Byte: %d\n", index, value)
	}

	// 9. Iterate over rune slice
	fmt.Println("\nIterating over Runes:")

	for index, value := range runes {
		fmt.Printf("Index: %d, Rune: %c\n", index, value)
	}

	// 10. Modify byte slice
	word := "Go"
	byteSlice := []byte(word)

	byteSlice[0] = 'N'

	fmt.Println("\nModified using byte slice:", string(byteSlice))

	// 11. Modify rune slice
	message := "கவி"
	runeSlice := []rune(message)

	runeSlice[0] = 'ந'

	fmt.Println("Modified using rune slice:", string(runeSlice))
}
