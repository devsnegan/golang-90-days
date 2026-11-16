package main

import "fmt"

type Employee interface {
	GetSalary() float64
}

type Developer struct {
	Salary float64
}

type Manager struct {
	Salary float64
}

func (d Developer) GetSalary() float64 {
	return d.Salary
}

func (m Manager) GetSalary() float64 {
	return m.Salary
}

func PrintSalary(e Employee) {
	fmt.Println("Salary:", e.GetSalary())
}

func main() {
	developer := Developer{Salary: 50000}
	manager := Manager{Salary: 80000}

	PrintSalary(developer)
	PrintSalary(manager)
}
