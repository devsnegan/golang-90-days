package main

import "fmt"

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultiFunction interface {
	Printer
	Scanner
}

type Machine struct{}

func (m Machine) Print() {
	fmt.Println("Printing document")
}

func (m Machine) Scan() {
	fmt.Println("Scanning document")
}

func main() {
	var device MultiFunction = Machine{}

	device.Print()
	device.Scan()
}