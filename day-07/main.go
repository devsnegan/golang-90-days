package main

import "fmt"

func main() {
	// 1. Arithmetic operators
	a, b := 10, 3

	fmt.Println("Arithmetic Operators")
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Integer Division:", a/b)
	fmt.Println("Remainder:", a%b)

	// 2. Comparison operators
	fmt.Println("\nComparison Operators")
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// 3. Logical operators
	isLoggedIn := true
	isAdmin := false

	fmt.Println("\nLogical Operators")
	fmt.Println("Logged in AND admin:", isLoggedIn && isAdmin)
	fmt.Println("Logged in OR admin:", isLoggedIn || isAdmin)
	fmt.Println("NOT logged in:", !isLoggedIn)

	// 4. Assignment operators
	number := 10
	number += 5
	fmt.Println("\nAssignment Operators")
	fmt.Println("After += 5:", number)

	number -= 3
	fmt.Println("After -= 3:", number)

	// 5. Increment and decrement
	count := 5
	count++
	fmt.Println("\nIncrement/Decrement")
	fmt.Println("After increment:", count)

	count--
	fmt.Println("After decrement:", count)

	// 6. Bitwise operators
	x, y := 6, 3 // Binary: 110 and 011

	fmt.Println("\nBitwise Operators")
	fmt.Println("x & y:", x&y)
	fmt.Println("x | y:", x|y)
	fmt.Println("x ^ y:", x^y)
	fmt.Println("x << 1:", x<<1)
	fmt.Println("x >> 1:", x>>1)

	// 7. Operator precedence
	result := 2 + 3*4
	fmt.Println("\nOperator Precedence")
	fmt.Println("2 + 3 * 4:", result)
}
