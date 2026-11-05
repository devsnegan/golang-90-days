package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Country string
}

type Employee struct {
	ID      int
	Name    string
	Address Address
}

func main() {
	employee := Employee{
		ID:   401,
		Name: "Kumar",
		Address: Address{
			Street:  "Anna Nagar",
			City:    "Chennai",
			State:   "Tamil Nadu",
			Country: "India",
		},
	}

	fmt.Println("Employee:", employee.Name)
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Country:", employee.Address.Country)
}
