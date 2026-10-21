package main

import (
	"fmt"
	"strconv"
)

func main() {
	pageText := "4"
	pageSize := 20

	page, err := strconv.Atoi(pageText)
	if err != nil || page < 1 {
		fmt.Println("Invalid page number")
		return
	}

	offset := (page - 1) * pageSize

	fmt.Println("Page:", page)
	fmt.Println("Page size:", pageSize)
	fmt.Println("Offset:", offset)
}
