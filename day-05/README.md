# Day 5 — Data Types in Go

**90 Days of Golang Mastery**  
**Author:** dev_snegan

## Overview

Today, I learned about Go's built-in data types and how to choose an appropriate type for different kinds of values.

Data types define what kind of value a variable can store and which operations can be performed on it.

## Learning Objectives

- Understand Go's built-in data types.
- Learn integer and floating-point types.
- Understand strings and Boolean values.
- Learn the difference between `byte` and `rune`.
- Understand zero values.
- Perform explicit type conversions.
- Understand string byte length versus rune count.

## 1. Integer Types

Go provides signed integer types:

- `int`
- `int8`
- `int16`
- `int32`
- `int64`

It also provides unsigned types such as `uint`, `uint8`, `uint16`, `uint32`, and `uint64`.

`int` is commonly used for integer operations and indexing. Fixed-width types are useful when a particular size is required.

## 2. Floating-Point Types

Go provides two floating-point types:

- `float32`
- `float64`

`float64` provides greater precision and is commonly used for floating-point calculations.

Floating-point types may introduce rounding errors. For financial calculations, consider integer minor units or a decimal arithmetic library.

## 3. Boolean Type

The `bool` type stores either `true` or `false`.

Example:

```go
isActive := true
isBlocked := false
```

Boolean values are useful for conditions, authentication checks, and account status.

## 4. String Type

The `string` type represents a sequence of bytes, commonly used for UTF-8 text.

Example:

```go
name := "Snegan"
fmt.Println(name)
fmt.Println(len(name))
```

The built-in `len()` function returns the number of bytes in a string, not necessarily its number of Unicode code points.

## 5. Byte and Rune

- `byte` is an alias for `uint8`.
- `rune` is an alias for `int32` and represents a Unicode code point.

Example:

```go
var letter byte = 'A'
var tamilLetter rune = 'அ'
```

Use `byte` when working with bytes and `rune` when iterating over Unicode code points.

## 6. Complex Types

Go supports `complex64` and `complex128` for complex numbers.

Example:

```go
number := complex(2, 3)

fmt.Println(number)
fmt.Println(real(number))
fmt.Println(imag(number))
```

Complex types are mainly useful in mathematical and scientific applications.

## 7. Zero Values

When a variable is declared without an explicit initial value, Go assigns its zero value.

| Type | Zero value |
|---|---|
| Integer | `0` |
| Floating-point | `0` |
| Boolean | `false` |
| String | `""` |
| Pointer | `nil` |
| Slice | `nil` |
| Map | `nil` |

## 8. Explicit Type Conversion

Go does not automatically convert incompatible numeric types during arithmetic.

Example:

```go
var a int = 10
var b float64 = 5.5

result := float64(a) + b
fmt.Println(result)
```

Output:

```text
15.5
```

The integer is explicitly converted to `float64` before the addition.

## Main Program

The main program in `main.go` demonstrates:

- Integer and floating-point types
- String and Boolean types
- Byte and rune values
- Complex numbers
- Zero values
- Explicit type conversion
- String byte length and rune count

## Exercises Completed

### Exercise 1 — Employee Profile

Created variables for an employee's name, ID, salary, and active status.

### Exercise 2 — Zero Values

Declared variables without explicit initialization and printed their default values.

### Exercise 3 — Numeric Conversion

Converted an integer to `float64` and performed arithmetic with a floating-point value.

### Exercise 4 — String Bytes and Runes

Compared byte length and rune count for a string containing a Tamil character.

### Exercise 5 — Fix a Type Mismatch

Corrected an invalid arithmetic operation by explicitly converting an integer to `float64`.

## Important Interview Questions

1. What is a data type in Go?
2. What is the difference between `int` and `int64`?
3. What is the difference between `float32` and `float64`?
4. What are zero values in Go?
5. What is the difference between `byte` and `rune`?
6. Does Go automatically convert numeric types?
7. What does `len()` return for a string?
8. Why should floating-point types be used carefully for money?

## Key Takeaways

- Go is statically typed.
- Each variable has a specific type.
- Go provides predictable zero values.
- Incompatible numeric types require explicit conversion.
- Strings are byte sequences; Unicode code points can be iterated over using `range`.
- Choosing suitable types helps prevent errors and improves program correctness.

## Completion Checklist

- [x] Learned the main built-in data types.
- [x] Practiced zero values.
- [x] Practiced explicit type conversion.
- [x] Explored byte and rune behavior.
- [x] Completed five exercises.
- [ ] Reviewed all interview questions.

**Next:** Day 6 — Type Conversion

---

*Learning Go step by step — @dev_snegan*
