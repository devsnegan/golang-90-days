
# Day 18 – Pointers in Go

## Topics Covered

- What is a pointer?
- Memory addresses
- Address operator (&)
- Dereference operator (*)
- Pointer declaration
- Nil pointers
- Modifying variables through pointers
- Passing pointers to functions
- Real-world bank deposit example

## Main Program

Run the main program:

```bash
go run ./day-18
```

## Exercises

| Exercise | Topic |
|---|---|
| 01 | Pointer Basics |
| 02 | Update Age Using a Pointer |
| 03 | Update Salary Using a Pointer |
| 04 | Swap Two Integers |
| 05 | Bank Deposit with Nil Validation |

## Important Concepts

- `&variable` gets the memory address.
- `*pointer` accesses the value stored at that address.
- A pointer's zero value is `nil`.
- Go passes arguments by value, including pointers.
- A copied pointer can still refer to the same underlying variable.

## Learning Goal

Understand how pointers work and how to modify
original variables through functions using pointers.
