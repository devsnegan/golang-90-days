# Day 17 – Named Return Values in Go

## Overview

Named return values allow us to declare names for function return variables directly in the function signature.

They make functions easier to understand when returning multiple related values.

## Syntax

```go
func calculate() (sum int, difference int) {
    sum = 30
    difference = 10

    return
}
```

## Normal vs Named Return Values

Normal return:

```go
func calculate(a, b int) (int, int) {
    return a + b, a - b
}
```

Named return:

```go
func calculate(a, b int) (sum int, difference int) {
    sum = a + b
    difference = a - b

    return
}
```

## Important Concepts

1. Named return variables are declared in the function signature.
2. They behave like local variables inside the function.
3. They are automatically initialized with their type's zero value.
4. A naked return returns the current named return variables.
5. Explicit return values can also be used.
6. Named returns can be useful when returning a result and an error.
7. Avoid excessive naked returns in complex functions.

## Zero Values

| Type | Zero Value |
|---|---|
| int | 0 |
| float64 | 0.0 |
| string | "" |
| bool | false |
| pointer | nil |
| slice | nil |
| map | nil |

## Real-World Applications

- Banking transactions
- Salary calculations
- Product pricing
- Student result calculations
- Backend service operations
- Error handling

## Exercises

- Exercise 01: Basic calculator.
- Exercise 02: Employee salary.
- Exercise 03: Rectangle calculator.
- Exercise 04: Student result.
- Exercise 05: Bank withdrawal.
- Exercise 06: Product discount.

## Key Takeaway

Named return values allow functions to declare and work with their return variables directly, improving readability when used appropriately.

**Day 17 completed – Named Return Values in Go.**

Next: Day 18 – Pointers in Go.