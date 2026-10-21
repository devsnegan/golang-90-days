package main

import (
	"fmt"
	"strconv"
)

func main() {
	text := "150"

	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Invalid integer:", err)
		return
	}

	fmt.Println("Converted number:", number)
}