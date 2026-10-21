package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("===== DAY 6: TYPE CONVERSION =====")

	// 1. Integer to float
	var age int = 25
	ageAsFloat := float64(age)

	fmt.Println("\n--- Integer to Float ---")
	fmt.Println("Original:", age)
	fmt.Println("Converted:", ageAsFloat)

	// 2. Float to integer
	price := 99.99
	priceAsInt := int(price)

	fmt.Println("\n--- Float to Integer ---")
	fmt.Println("Original:", price)
	fmt.Println("Converted:", priceAsInt)

	// 3. Integer type conversion
	var number int = 100
	numberAsInt64 := int64(number)

	fmt.Println("\n--- Integer Conversion ---")
	fmt.Println("int:", number)
	fmt.Println("int64:", numberAsInt64)

	// 4. Integer to string
	userID := 2026
	userIDText := strconv.Itoa(userID)

	fmt.Println("\n--- Integer to String ---")
	fmt.Println("Number:", userID)
	fmt.Println("String:", userIDText)

	// 5. String to integer
	ageText := "25"
	parsedAge, err := strconv.Atoi(ageText)

	fmt.Println("\n--- String to Integer ---")
	if err != nil {
		fmt.Println("Conversion failed:", err)
	} else {
		fmt.Println("Parsed age:", parsedAge)
	}

	// 6. String to float
	priceText := "149.50"
	parsedPrice, err := strconv.ParseFloat(priceText, 64)

	fmt.Println("\n--- String to Float ---")
	if err != nil {
		fmt.Println("Conversion failed:", err)
	} else {
		fmt.Println("Parsed price:", parsedPrice)
	}

	// 7. Invalid input handling
	invalidText := "hello"
	_, err = strconv.Atoi(invalidText)

	fmt.Println("\n--- Invalid Input ---")
	if err != nil {
		fmt.Println("Cannot convert:", invalidText)
		fmt.Println("Error:", err)
	}

	// 8. Backend pagination example
	pageText := "3"
	page, err := strconv.Atoi(pageText)

	fmt.Println("\n--- Pagination Example ---")
	if err != nil || page < 1 {
		fmt.Println("Invalid page number")
	} else {
		pageSize := 10
		offset := (page - 1) * pageSize

		fmt.Println("Page:", page)
		fmt.Println("Page size:", pageSize)
		fmt.Println("Offset:", offset)
	}
}