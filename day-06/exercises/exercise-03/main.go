package main

import (
	"fmt"
	"strconv"
)

func main() {
	year := 2026
	yearText := strconv.Itoa(year)

	fmt.Println("Year:", year)
	fmt.Println("Year as string:", yearText)
}