package main

import "fmt"

func main() {
	// 1. Create and initialize a map
	ages := map[string]int{
		"Arun":  25,
		"Bala":  28,
		"Kumar": 30,
	}
	fmt.Println("Student ages:", ages)

	// 2. Access values
	fmt.Println("Arun's age:", ages["Arun"])

	// 3. Add a new entry
	ages["Snegan"] = 22
	fmt.Println("After adding Snegan:", ages)

	// 4. Update an entry
	ages["Arun"] = 26
	fmt.Println("After updating Arun:", ages)

	// 5. Check whether a key exists
	age, exists := ages["Snegan"]
	if exists {
		fmt.Println("Snegan's age:", age)
	} else {
		fmt.Println("Snegan not found")
	}

	// 6. Check a missing key
	age, exists = ages["Ravi"]
	fmt.Println("Missing key value:", age)
	fmt.Println("Ravi exists:", exists)

	// 7. Iterate through a map
	fmt.Println("All students:")
	for name, studentAge := range ages {
		fmt.Printf("%s: %d\n", name, studentAge)
	}

	// 8. Count entries
	fmt.Println("Total students:", len(ages))

	// 9. Delete an entry
	delete(ages, "Bala")
	fmt.Println("After deleting Bala:", len(ages))

	// 10. Create a map using make
	scores := make(map[string]int)
	scores["Arun"] = 80
	scores["Bala"] = 90
	fmt.Println("Scores:", scores)

	// 11. Demonstrate map assignment
	first := map[string]int{"Arun": 80}
	second := first
	second["Arun"] = 100

	fmt.Println("First map:", first)
	fmt.Println("Second map:", second)

	// 12. Read a nil map safely
	var nilMap map[string]int
	fmt.Println("Nil map:", nilMap)
	fmt.Println("Nil map length:", len(nilMap))
	fmt.Println("Nil map lookup:", nilMap["Arun"])

	// 13. Add an entry to a nil map safely
	nilMap = make(map[string]int)
	nilMap["Arun"] = 95
	fmt.Println("Initialized nil map:", nilMap)
}
