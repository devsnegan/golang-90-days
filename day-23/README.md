# Day 23 - Empty Interface and any

## Topic

Empty Interface (`interface{}`) and `any`

## What I Learned

- What is an empty interface?
- Why do we need an empty interface?
- Difference between interface{} and any.
- How any accepts different data types.
- Storing different values in an interface variable.
- Interface zero value and nil.
- Dynamic type and dynamic value.
- Practical backend logging example.
- Limitations of using any.
- Difference between any and generics.

## Main Program

Created a PrintValue function that accepts any type.

Demonstrated:
- Integer
- String
- Float
- Boolean
- Struct
- Nil interface value

## Exercises

1. Print Different Values
2. Store Different Values
3. Employee Information
4. Check Nil
5. Flexible Logger

## Important Concepts

- any is an alias for interface{}.
- Every Go type satisfies the empty interface.
- The zero value of any is nil.
- An interface can hold different concrete types.
- any does not allow arbitrary operations.
- Prefer specific types when possible.

## Go Version

Go 1.18 or later is required for the any alias.

## Commands

Run main program:

go run ./day-23

Run exercises individually:

go run ./day-23/exercises/exercise-01
go run ./day-23/exercises/exercise-02
go run ./day-23/exercises/exercise-03
go run ./day-23/exercises/exercise-04
go run ./day-23/exercises/exercise-05