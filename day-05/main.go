package main

import "fmt"

func main() {
	// 1. Integer types
	var age int = 25
	var population int64 = 1000000

	// 2. Floating-point types
	var temperature float32 = 36.5
	var salary float64 = 45000.50

	// 3. Boolean and string types
	var name string = "Snegan"
	var isDeveloper bool = true

	// 4. Byte and rune types
	var letter byte = 'A'
	var tamilLetter rune = 'அ'

	// 5. Complex type
	var number complex128 = complex(2, 3)

	fmt.Println("===== DAY 5: GO DATA TYPES =====")

	fmt.Println("\n--- Integer Types ---")
	fmt.Println("Age:", age)
	fmt.Println("Population:", population)

	fmt.Println("\n--- Floating-Point Types ---")
	fmt.Println("Temperature:", temperature)
	fmt.Println("Salary:", salary)

	fmt.Println("\n--- String and Boolean ---")
	fmt.Println("Name:", name)
	fmt.Println("Is Developer:", isDeveloper)

	fmt.Println("\n--- Byte and Rune ---")
	fmt.Println("Byte value:", letter)
	fmt.Printf("Byte as character: %c\n", letter)
	fmt.Println("Rune code point:", tamilLetter)
	fmt.Printf("Rune as character: %c\n", tamilLetter)

	fmt.Println("\n--- Complex Type ---")
	fmt.Println("Complex number:", number)
	fmt.Println("Real part:", real(number))
	fmt.Println("Imaginary part:", imag(number))

	// 6. Zero values
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("\n--- Zero Values ---")
	fmt.Println("Integer:", defaultInt)
	fmt.Println("Float:", defaultFloat)
	fmt.Println("Boolean:", defaultBool)
	fmt.Printf("String: %q\n", defaultString)

	// 7. Explicit type conversion
	var a int = 10
	var b float64 = 5.5
	result := float64(a) + b

	fmt.Println("\n--- Type Conversion ---")
	fmt.Println("Result:", result)

	// 8. String bytes versus runes
	text := "Go అ"

	fmt.Println("\n--- String Length ---")
	fmt.Println("Text:", text)
	fmt.Println("Byte length:", len(text))
	fmt.Println("Rune count:", len([]rune(text)))

	fmt.Println("\n--- Iterating Over Runes ---")
	for _, r := range text {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}
