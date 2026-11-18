package main

import "fmt"

func PrintValue(value any) {
	fmt.Println(value)
}

func main() {
	PrintValue(100)
	PrintValue("Golang")
	PrintValue(25.5)
	PrintValue(true)
}