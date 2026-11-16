package main

import "fmt"

// Interface declaration
type Animal interface {
	Sound()
}

// Concrete types
type Dog struct{}
type Cat struct{}
type Cow struct{}

// Dog implements Animal
func (d Dog) Sound() {
	fmt.Println("Dog barks")
}

// Cat implements Animal
func (c Cat) Sound() {
	fmt.Println("Cat meows")
}

// Cow implements Animal
func (c Cow) Sound() {
	fmt.Println("Cow moos")
}

// Function accepts any type satisfying Animal
func MakeSound(a Animal) {
	a.Sound()
}

// Interface with a pointer receiver implementation
type Printer interface {
	Print()
}

type Document struct {
	Name string
}

func (d *Document) Print() {
	fmt.Println("Printing document:", d.Name)
}

func main() {
	fmt.Println("=== Multiple Types, One Interface ===")

	MakeSound(Dog{})
	MakeSound(Cat{})
	MakeSound(Cow{})

	fmt.Println("\n=== Interface Variable ===")

	var animal Animal

	animal = Dog{}
	fmt.Println("Dog assigned to Animal interface")
	animal.Sound()

	animal = Cat{}
	fmt.Println("Cat assigned to Animal interface")
	animal.Sound()

	fmt.Println("\n=== Nil Interface ===")

	var emptyAnimal Animal
	fmt.Println("Is interface nil?", emptyAnimal == nil)

	fmt.Println("\n=== Pointer Receiver ===")

	document := Document{Name: "Go Interfaces"}
	var printer Printer

	printer = &document
	printer.Print()
}
