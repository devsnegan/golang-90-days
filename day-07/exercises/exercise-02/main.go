package main

import "fmt"

func main() {
	age := 22
	minimumAge := 18

	fmt.Println("Eligible:", age >= minimumAge)
	fmt.Println("Exactly 18:", age == 18)
	fmt.Println("Below 18:", age < 18)
	fmt.Println("Not 18:", age != 18)
}