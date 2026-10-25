# Day 9: for Loops in Go

## Overview

Today, I learned how to repeat instructions using Go's `for` loop and how to iterate through collections using `range`.

## Topics Covered

- Traditional `for` loops
- Initialization, condition, and post statement
- Condition-only loops
- Infinite loops
- `break` and `continue`
- Iterating over slices with `range`
- Index and value in `range`
- Ignoring values with `_`
- Nested loops
- Calculating totals with loops

## Project Structure

```text
day-09/
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
go run ./day-09
```

Run an individual exercise:

```bash
go run ./day-09/exercises/exercise-01
```

## Key Learnings

- Go uses `for` as its only loop keyword.
- The loop condition determines when the loop stops.
- `range` visits each item in a collection.
- The index starts at zero for slices and arrays.
- `_` allows an unused value to be ignored.
- `continue` skips the current iteration.
- `break` exits the loop.
- Loops help process collections and perform repeated calculations.

## Practical Applications

Loops are useful for:

- Processing orders and transactions
- Calculating totals
- Iterating through API results
- Processing records returned from a database
- Running background jobs
- Applying business rules to collections

## Progress

- [x] Day 9: for Loops
- [ ] Day 10: Arrays

Repository: https://github.com/devsnegan/golang-90-days
