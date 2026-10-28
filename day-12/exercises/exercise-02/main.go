package main

import "fmt"

func main() {
	scores := map[string]int{
		"Arun": 80,
	}

	scores["Bala"] = 90
	scores["Arun"] = 95

	fmt.Println("Scores:", scores)
}
