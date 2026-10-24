# Day 8: if, else, and switch in Go

## Overview

Today, I learned how to control program execution using conditional statements in Go.

## Topics Covered

- `if` statements
- `if-else` statements
- `else if` chains
- Nested `if` statements
- Variables declared in `if` initializers
- `switch` statements
- Multiple values in a `case`
- Expressionless `switch`
- `default` and `fallthrough`

## Project Structure

```text
day-08/
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
go run ./day-08
```

Run an individual exercise:

```bash
go run ./day-08/exercises/exercise-01
```

## Key Learnings

- `if` executes code when a condition is true.
- `else` handles the alternative path.
- `else if` checks additional conditions.
- Nested conditions allow decisions to depend on earlier checks.
- `switch` is useful when selecting among multiple alternatives.
- Go automatically exits a matched `switch` case unless `fallthrough` is used.
- Conditions must be ordered carefully when checking ranges.

## Practical Applications

Conditional statements are useful for:

- User registration validation
- Login and authorization decisions
- Account status checks
- Payment and balance validation
- HTTP status handling
- Business rule implementation

## Progress

- [x] Day 8: if, else, and switch
- [ ] Day 9: for Loops

Repository: https://github.com/devsnegan/golang-90-days
