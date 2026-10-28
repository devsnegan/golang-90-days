package main

import "fmt"

func main() {
	ages := map[string]int{
		"Arun": 25,
		"Bala": 0,
	}

	age, exists := ages["Bala"]
	fmt.Println("Bala:", age, "Exists:", exists)

	age, exists = ages["Snegan"]
	fmt.Println("Snegan:", age, "Exists:", exists)
}
