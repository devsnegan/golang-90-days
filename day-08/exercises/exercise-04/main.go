package main

import "fmt"

func main() {
	score := 82

	if score < 0 || score > 100 {
		fmt.Println("Invalid score")
	} else if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else if score >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
}