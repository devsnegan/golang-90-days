
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("Dog says: Woof!")
}

func (c Cat) Speak() {
	fmt.Println("Cat says: Meow!")
}

func main() {
	var speaker Speaker

	speaker = Dog{}
	speaker.Speak()

	speaker = Cat{}
	speaker.Speak()
}
