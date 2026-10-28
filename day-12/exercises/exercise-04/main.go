package main

import "fmt"

func main() {
	students := map[string]int{
		"Arun":  80,
		"Bala":  90,
		"Kumar": 85,
	}

	delete(students, "Bala")

	fmt.Println("Students:", students)
	fmt.Println("Remaining entries:", len(students))
}
