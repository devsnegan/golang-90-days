# Day 10: Arrays in Go

## Overview

Today I learned about **Arrays in Go**. An array is a collection of elements of the same data type with a fixed length.

Arrays help us store and process multiple values using a single variable.

## Topics Covered

- What are arrays?
- Declaring arrays
- Initializing arrays
- Zero values in arrays
- Accessing elements using indexes
- Updating array elements
- Finding array length using `len()`
- Iterating using `for` loops
- Iterating using `range`
- Calculating the sum and average
- Finding the largest element
- Counting even and odd numbers
- Reversing an array
- Copying arrays
- Understanding array types and fixed lengths

## 1. Declaring an Array

```go
var numbers [5]int

fmt.Println(numbers)
```

Output:

```text
[0 0 0 0 0]
```

An integer array contains zeros when its elements have not been initialized.

## 2. Initializing an Array

```go
numbers := [5]int{10, 20, 30, 40, 50}

fmt.Println(numbers)
```

Output:

```text
[10 20 30 40 50]
```

Go also supports automatic length inference:

```go
numbers := [...]int{10, 20, 30}
```

## 3. Accessing Array Elements

Array indexes start at zero.

```go
numbers := [3]int{10, 20, 30}

fmt.Println(numbers[0])
fmt.Println(numbers[1])
fmt.Println(numbers[2])
```

Output:

```text
10
20
30
```

## 4. Updating an Array Element

```go
numbers := [3]int{10, 20, 30}

numbers[1] = 100

fmt.Println(numbers)
```

Output:

```text
[10 100 30]
```

## 5. Finding Array Length

Use the built-in `len()` function.

```go
numbers := [4]int{10, 20, 30, 40}

fmt.Println(len(numbers))
```

Output:

```text
4
```

## 6. Iterating Using a For Loop

```go
numbers := [3]int{10, 20, 30}

for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

Output:

```text
10
20
30
```

## 7. Iterating Using Range

```go
numbers := [3]int{10, 20, 30}

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

Use `_` when you do not need the index:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

## 8. Calculating the Sum

```go
numbers := [5]int{10, 20, 30, 40, 50}
sum := 0

for _, value := range numbers {
    sum += value
}

fmt.Println("Sum:", sum)
```

Output:

```text
Sum: 150
```

## 9. Important Characteristics of Arrays

- Arrays have a fixed length.
- All elements have the same data type.
- Indexes start at zero.
- `len()` returns the number of elements.
- Array length is part of its type. For example, `[3]int` and `[4]int` are different types.
- Assigning an array to another variable copies its elements.
- Accessing an index outside the array bounds causes a runtime panic.

## 10. Arrays Are Copied by Value

```go
original := [3]int{10, 20, 30}
copied := original

copied[0] = 100

fmt.Println("Original:", original)
fmt.Println("Copied:", copied)
```

Output:

```text
Original: [10 20 30]
Copied: [100 20 30]
```

Changing the copied array does not change the original array.

## 11. Arrays vs Slices

| Arrays | Slices |
|---|---|
| Fixed length | Flexible length |
| Length is part of the type | Length can change as elements are appended |
| Declared using `[N]Type` | Declared using `[]Type` |
| Useful when size is fixed | Commonly used for collections of changing size |

## Exercises Completed

1. Declare and print an array.
2. Calculate the sum of array elements.
3. Find the largest number.
4. Count even and odd numbers.
5. Reverse an array.
6. Copy an array and modify the copy.

## Commands Used

Run the main lesson:

```bash
go run ./day-10
```

Run an exercise:

```bash
go run ./day-10/exercises/exercise-01
```

Format the code:

```bash
gofmt -w day-10
```

Run tests across the repository:

```bash
go test ./...
```

## Key Takeaways

- Arrays store multiple values of the same type.
- Array length is fixed.
- Use indexes to access and update elements.
- Use `for` or `range` to iterate over elements.
- Array assignments copy the array's values.
- Slices provide a more flexible way to work with collections in Go.

## Next Day

**Day 11: Slices in Go**

Next, I will learn how slices work, how they relate to arrays, and how to use `append()`, `len()`, and `cap()`.
