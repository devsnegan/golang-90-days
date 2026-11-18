package main

import "fmt"

func LogInfo(value any) {
	fmt.Println("[INFO]:", value)
}

func main() {
	LogInfo("Application started")
	LogInfo(101)
	LogInfo(45000.50)
	LogInfo(true)
}
