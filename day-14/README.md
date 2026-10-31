# Day 14 - Functions in Go

## Overview

Today I learned how to create and use functions in Go. Functions help organize code into smaller, reusable blocks.

## Topics Covered

1. What is a function?
2. Why functions are needed
3. Function declaration
4. Function calling
5. Function execution flow
6. Function parameters
7. Parameters vs arguments
8. Multiple parameters
9. Return values
10. Different return types
11. Functions calling other functions
12. Variable scope

## Basic Function

```go
func greet() {
    fmt.Println("Hello Go")
}
```

## Function with Parameters

```go
func greet(name string) {
    fmt.Println("Hello", name)
}
```

## Function with Multiple Parameters

```go
func add(a, b int) {
    fmt.Println(a + b)
}
```

## Function with Return Value

```go
func multiply(a, b int) int {
    return a * b
}
```

## Function Returning Boolean

```go
func isEven(number int) bool {
    return number%2 == 0
}
```

## Parameter vs Argument

| Parameter | Argument |
|---|---|
| Declared in function definition | Passed during function call |
| Receives a value | Supplies a value |
| `name string` | `"Snegan"` |

## Practical Applications

- Reusable business logic
- Banking operations
- Calculating prices
- Validating user input
- Processing API requests
- Separating application responsibilities

## Exercises Completed

- [x] Exercise 01 - Simple Greeting
- [x] Exercise 02 - Function with Parameter
- [x] Exercise 03 - Addition Function
- [x] Exercise 04 - Even or Odd
- [x] Exercise 05 - Rectangle Area
- [x] Exercise 06 - Banking Deposit

## Key Takeaways

- Functions are declared using `func`.
- Parameters receive input values.
- Arguments are passed when calling functions.
- Return values send results back to the caller.
- Functions help make code reusable and maintainable.
- Variables declared inside functions have their own scope.

## Next Day

Day 15 - Variadic Functions in Go