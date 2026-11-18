package main

import "fmt"

func main() {
	var data any

	fmt.Println("Before assignment:", data == nil)

	data = "Go Programming"

	fmt.Println("After assignment:", data == nil)
}