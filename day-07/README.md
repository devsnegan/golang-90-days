# Day 7: Operators in Go

## Overview

Today, I learned about operators in Go and how they are used to perform calculations, compare values, evaluate conditions, and manipulate bits.

## Topics Covered

- Arithmetic operators: `+`, `-`, `*`, `/`, `%`
- Comparison operators: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical operators: `&&`, `||`, `!`
- Assignment operators: `=`, `+=`, `-=`, `*=`, `/=`, `%=`
- Increment and decrement: `++`, `--`
- Bitwise operators: `&`, `|`, `^`, `&^`
- Bit shifting: `<<`, `>>`
- Operator precedence
- Practical examples of operators in backend development

## Project Structure

```text
day-07/
├── main.go
├── README.md
└── exercises/
    ├── exercise-01/main.go
    ├── exercise-02/main.go
    ├── exercise-03/main.go
    ├── exercise-04/main.go
    ├── exercise-05/main.go
    └── exercise-06/main.go
```

## How to Run

Run the main program:

```bash
go run ./day-07
```

Run an individual exercise:

```bash
go run ./day-07/exercises/exercise-01
```

## Key Learnings

- Integer division discards the fractional part.
- The modulo operator returns the remainder.
- Comparison operators return a boolean.
- Logical operators combine or negate conditions.
- Go requires explicit parentheses when grouping expressions for clarity.
- Go supports postfix `++` and `--` as statements, not as expressions.
- Bitwise operators manipulate individual bits.
- Operator precedence determines the order of evaluation.

## Practical Application

Operators are used in backend development for:

- Calculating prices and discounts
- Validating user permissions
- Checking pagination values
- Applying business rules
- Managing counters and limits

## Progress

- [x] Day 7: Operators in Go\
Day 8: if / else / switch

Repository: [golang-90-days](https://github.com/devsnegan/golang-90-days)
