
# Day 21 - Interfaces in Go

## Overview

Today I learned about interfaces in Go and how different types can
implement the same behavior.

## Topics Covered

1. What is an interface?
2. Why interfaces are needed
3. Declaring interfaces
4. Implicit interface implementation
5. Interface values
6. Nil interfaces
7. Value receiver and pointer receiver with interfaces
8. Real-world backend use cases

## Main Example

The main program demonstrates:

- Animal interface
- Dog, Cat, and Cow implementations
- Passing interfaces to functions
- Interface variables
- Nil interface
- Pointer receiver implementation

## Exercises

### Exercise 01 - Shape Interface
Implemented Shape using Rectangle and Circle.

### Exercise 02 - Speaker Interface
Implemented Speaker using Dog and Cat.

### Exercise 03 - Payment Method Interface
Implemented PaymentMethod using UPI and CreditCard.

### Exercise 04 - Notifier Interface
Implemented Notifier using EmailNotifier and SMSNotifier.

### Exercise 05 - Employee Salary Interface
Implemented Employee using Developer and Manager.

## Important Concepts

- An interface defines method requirements.
- Go uses implicit interface implementation.
- A type satisfies an interface by implementing all required methods.
- Multiple types can satisfy the same interface.
- Interfaces help build reusable and loosely coupled code.
- Pointer receivers affect interface satisfaction.

## Commands

Run the main program:

    go run ./day-21

Run an exercise:

    go run ./day-21/exercises/exercise-01

Format code:

    gofmt -w day-21

Run tests:

    go test ./...

## Learning Progress

Phase 2 - Go Core

Day 21 of 90 completed.

Next: Day 22 - Interface Composition
