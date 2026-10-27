# Day 11: Slices in Go

## Overview

Today I learned about **Slices in Go**. A slice is a flexible data structure used to work with a sequence of elements of the same type.

Unlike arrays, slices can grow when new elements are appended.

## Topics Covered

- What are slices?
- Declaring and initializing slices
- Accessing and updating elements
- Using `append()`
- Understanding `len()` and `cap()`
- Creating slices using `make()`
- Slicing expressions
- Iterating using `range`
- Sharing an underlying array
- Copying slices using `copy()`
- Nil slices vs empty slices
- Calculating sum and average
- Finding the largest element
- Counting even and odd numbers

## 1. Creating a Slice

```go
numbers := []int{10, 20, 30, 40, 50}

fmt.Println(numbers)
```

Output:

```text
[10 20 30 40 50]
```

## 2. Appending Elements

```go
numbers := []int{10, 20, 30}
numbers = append(numbers, 40, 50)

fmt.Println(numbers)
```

Output:

```text
[10 20 30 40 50]
```

Always assign the result of `append()` back to a slice variable.

## 3. Length and Capacity

- `len()` returns the number of elements.
- `cap()` returns the capacity available from the slice's starting position.

```go
numbers := make([]int, 3, 5)

fmt.Println(len(numbers))
fmt.Println(cap(numbers))
```

Output:

```text
3
5
```

## 4. Slicing an Existing Slice

```go
numbers := []int{10, 20, 30, 40, 50}

part := numbers[1:4]

fmt.Println(part)
```

Output:

```text
[20 30 40]
```

The start index is included, but the end index is excluded.

## 5. Iterating Using Range

```go
numbers := []int{10, 20, 30}

for index, value := range numbers {
    fmt.Println(index, value)
}
```

Output:

```text
0 10
1 20
2 30
```

Use `_` when the index is not needed.

## 6. Copying a Slice

```go
original := []int{10, 20, 30}

copied := make([]int, len(original))
copy(copied, original)

copied[0] = 100

fmt.Println(original)
fmt.Println(copied)
```

Output:

```text
[10 20 30]
[100 20 30]
```

The copied slice has its own underlying array in this example.

## 7. Important Concepts

1. Slices are more flexible than fixed-length arrays.
2. Slice indexes start at zero.
3. `append()` adds elements and returns a slice.
4. `len()` gives the current number of elements.
5. `cap()` gives the capacity from the slice's starting position.
6. `make()` can specify the initial length and capacity.
7. Slicing an array or slice can share the underlying array.
8. `copy()` copies elements into a destination slice.
9. Nil and empty slices are different.
10. Capacity growth is controlled by Go and should not be assumed to follow an exact pattern.

## Exercises Completed

- Exercise 01: Create and access a slice
- Exercise 02: Append elements
- Exercise 03: Calculate sum and average
- Exercise 04: Find the largest number
- Exercise 05: Count even and odd numbers
- Exercise 06: Copy a slice independently

## Commands Used

Run the main lesson:

```bash
go run ./day-11
```

Run an exercise:

```bash
go run ./day-11/exercises/exercise-01
```

Format the Go files:

```bash
gofmt -w day-11
```

Run tests and check packages:

```bash
go test ./...
```

## Key Takeaways

Slices are one of the most commonly used data structures in Go. Understanding `append()`, `len()`, `cap()`, `make()`, and `copy()` is essential for writing effective Go programs.

## Next Day

**Day 12: Maps in Go**
