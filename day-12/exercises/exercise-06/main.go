package main

import "fmt"

func main() {
	words := []string{
		"go", "java", "go", "python", "go", "java",
	}

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	fmt.Println("Word frequency:")

	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
