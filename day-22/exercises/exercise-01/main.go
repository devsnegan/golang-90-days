package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() {
	fmt.Println("Reading...")
}

func (f File) Write() {
	fmt.Println("Writing...")
}

func Process(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	file := File{}
	Process(file)
}