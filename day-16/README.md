# Day 16 – Multiple Return Values in Go

## Overview

Go functions can return multiple values. This feature is useful when a function needs to return a result along with additional information, such as an error or status.

## Syntax

```go
func functionName() (type1, type2) {
    return value1, value2
}
```

## Example 1 – Return multiple values

```go
func calculate(a, b int) (int, int) {
    return a + b, a - b
}
```

## Example 2 – Return a result and an error

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }

    return a / b, nil
}
```

## Important Concepts

1. A Go function can return multiple values.
2. Return types are declared inside parentheses.
3. The caller must receive the expected number of values.
4. The blank identifier `_` ignores an unwanted return value.
5. `error` is commonly returned alongside a result.
6. `nil` indicates that no error occurred.
7. Always check errors before using the returned result.

## Real-World Applications

- Database queries
- User lookup
- API responses
- File operations
- Banking transactions
- Input validation

## Exercises

- Exercise 01: Addition and subtraction.
- Exercise 02: Multiplication and division.
- Exercise 03: Find minimum and maximum.
- Exercise 04: Check even or odd.
- Exercise 05: User lookup with error handling.
- Exercise 06: Bank withdrawal with balance validation.

## Key Takeaway

Multiple return values make Go functions more expressive and provide a simple way to return results together with errors.

**Day 16 completed – Multiple Return Values in Go.**

Next: Day 17 – Named Return Values.