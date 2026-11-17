# Day 22 - Interface Composition

## Topic
Interface Composition in Go

## What I Learned

- What is interface composition?
- Why interface composition is useful.
- Embedding one interface inside another.
- Combining multiple interfaces.
- Implicit interface implementation.
- Interface composition vs struct embedding.
- Real-world backend repository example.

## Main Example

Created Reader and Writer interfaces.

Composed both interfaces into ReadWriter.

Implemented both methods using the File struct.

## Exercises

1. Reader and Writer
2. Printer and Scanner
3. Vehicle Capabilities
4. Authentication and Authorization
5. User Repository

## Important Concepts

- An interface can embed multiple interfaces.
- A type must implement all required methods.
- Go does not require an implements keyword.
- Small interfaces improve code reusability.
- Interface composition is useful in backend architecture.

## Commands

Run main program:

go run ./day-22

Run all exercises individually:

go run ./day-22/exercises/exercise-01
go run ./day-22/exercises/exercise-02
go run ./day-22/exercises/exercise-03
go run ./day-22/exercises/exercise-04
go run ./day-22/exercises/exercise-05