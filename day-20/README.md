# Day 15 – Variadic Functions in Go

## Overview

A variadic function is a function that accepts zero or more arguments of the same declared type.

It is useful when the number of arguments is not known in advance.

## Syntax

```go
func functionName(values ...int) {
    // function body
}
```

The `...` syntax allows the function to receive multiple arguments.

Inside the function, the variadic parameter behaves like a slice.

## Example

```go
package main

import "fmt"

func add(numbers ...int) int {
    total := 0

    for _, number := range numbers {
        total += number
    }

    return total
}

func main() {
    fmt.Println(add(10, 20, 30))
}
```

Output:

```text
60
```

## Important Concepts

1. A variadic function can accept zero or more arguments.
2. A variadic parameter behaves like a slice inside the function.
3. A variadic parameter must be the last parameter.
4. Only one variadic parameter is allowed in a function.
5. Use `slice...` to pass a slice to a variadic function.
6. Go's `fmt.Println` is an example of a variadic function.
7. `...any` allows different types of arguments.

## Slice Expansion

```go
numbers := []int{10, 20, 30}

result := add(numbers...)
```

The `...` expands the slice into individual arguments.

## Normal and Variadic Parameters

```go
func greet(message string, names ...string) {
    for _, name := range names {
        fmt.Println(message, name)
    }
}
```

## Exercises

- Exercise 01: Print multiple numbers.
- Exercise 02: Calculate the sum.
- Exercise 03: Find the largest number.
- Exercise 04: Calculate the average.
- Exercise 05: Greet multiple users.
- Exercise 06: Calculate shopping cart total.

## Real-World Applications

- Logging functions
- Formatting functions
- Shopping cart calculations
- Sending notifications to multiple users
- Utility functions that accept flexible inputs

## Key Takeaway

Variadic functions make Go functions flexible by allowing them to accept a variable number of arguments.

**Day 15 completed – Phase 1: Go Fundamentals completed!**

Next: Phase 1 Mini Project – CLI Student Management System.