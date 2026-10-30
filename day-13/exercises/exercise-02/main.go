
package main

import "fmt"

func main() {
	text := "Golang"

	bytes := []byte(text)

	fmt.Println("Original String:", text)
	fmt.Println("Byte Slice:", bytes)
	fmt.Println("Byte Count:", len(bytes))

	newText := string(bytes)

	fmt.Println("Converted String:", newText)
}
