package main

import "fmt"

// First interface
type Reader interface {
	Read()
}

// Second interface
type Writer interface {
	Write()
}

// Composed interface
type ReadWriter interface {
	Reader
	Writer
}

// Struct implementing both interfaces
type File struct {
	Name string
}

// Implement Reader
func (f File) Read() {
	fmt.Println("Reading file:", f.Name)
}

// Implement Writer
func (f File) Write() {
	fmt.Println("Writing file:", f.Name)
}

// Function accepting composed interface
func ProcessFile(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	file := File{
		Name: "golang.txt",
	}

	fmt.Println("Interface Composition Example")
	ProcessFile(file)
}
