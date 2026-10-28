package main

import "fmt"

func main() {
	studentMarks := map[string]int{
		"Arun":  80,
		"Bala":  90,
		"Kumar": 85,
	}

	fmt.Println("Student marks:", studentMarks)
	fmt.Println("Arun's marks:", studentMarks["Arun"])
	fmt.Println("Total students:", len(studentMarks))
}
