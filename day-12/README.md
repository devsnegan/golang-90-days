# Day 12: Maps in Go

## Overview

Today I learned about **Maps in Go**. A map stores data as key-value pairs, allowing values to be accessed using their keys.

Maps are useful for lookups, counting occurrences, storing configuration, and managing records.

## Topics Covered

- Understanding maps
- Map syntax and data types
- Creating maps using literals
- Creating maps using `make()`
- Accessing map values
- Adding and updating entries
- Deleting entries using `delete()`
- Checking key existence using comma-ok
- Iterating with `range`
- Counting entries using `len()`
- Nil maps
- Map assignment behavior
- Word-frequency counting

## 1. Creating a Map

```go
ages := map[string]int{
    "Arun": 25,
    "Bala": 28,
}

fmt.Println(ages)
```

A `map[string]int` stores string keys and integer values.

## 2. Adding and Updating Entries

```go
ages := make(map[string]int)

ages["Arun"] = 25
ages["Bala"] = 28

ages["Arun"] = 26
```

Assigning a value to a new key adds an entry. Assigning to an existing key updates it.

## 3. Accessing Values

```go
fmt.Println(ages["Arun"])
```

If the key does not exist, Go returns the zero value of the map's value type.

## 4. Checking Key Existence

```go
age, exists := ages["Arun"]

if exists {
    fmt.Println("Age:", age)
} else {
    fmt.Println("Student not found")
}
```

The boolean `exists` indicates whether the key is present.

## 5. Deleting Entries

```go
delete(ages, "Bala")
```

The `delete()` function removes the specified entry.

## 6. Iterating Over a Map

```go
for name, age := range ages {
    fmt.Println(name, age)
}
```

Map iteration order is unspecified and should not be relied upon.

## 7. Counting Entries

```go
fmt.Println(len(ages))
```

The `len()` function returns the number of entries in the map.

## 8. Nil Maps

```go
var scores map[string]int
```

A nil map can be read, but assigning an entry to it causes a runtime panic. Initialize it with `make()` or a map literal before writing.

## 9. Important Concepts

1. Maps store key-value pairs.
2. Keys must be comparable types.
3. Maps can be created using `make()` or map literals.
4. Assignment can add or update an entry.
5. Use comma-ok to check whether a key exists.
6. Use `delete()` to remove an entry.
7. Map iteration order is unspecified.
8. Assigning a map to another variable does not create an independent copy.
9. Ordinary maps are not safe for concurrent reads and writes without synchronization.

## Exercises Completed

- Exercise 01: Create and access a map
- Exercise 02: Add and update map entries
- Exercise 03: Check whether a key exists
- Exercise 04: Delete a map entry
- Exercise 05: Iterate over a map
- Exercise 06: Count word frequency

## Commands Used

Run the main lesson:

```bash
go run ./day-12
```

Run an exercise:

```bash
go run ./day-12/exercises/exercise-01
```

Format the Go files:

```bash
gofmt -w day-12
```

Run tests and check packages:

```bash
go test ./...
```

## Key Takeaways

Maps provide convenient key-based lookups and updates. Understanding initialization, key existence, deletion, iteration, and map assignment is essential for writing reliable Go programs.

## Next Day

**Day 13: Strings & Runes in Go**
