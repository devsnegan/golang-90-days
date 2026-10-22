package main

import "fmt"

func main() {
	a := 12 // Binary: 1100
	b := 10 // Binary: 1010

	fmt.Println("AND:", a&b)
	fmt.Println("OR:", a|b)
	fmt.Println("XOR:", a^b)
	fmt.Println("Left Shift:", a<<1)
	fmt.Println("Right Shift:", a>>1)
}