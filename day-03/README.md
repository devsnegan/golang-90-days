# Day 3 — Program Structure in Go

**Phase:** 1 — Go Fundamentals  
**Day:** 3 of 90  
**Topic:** Program Structure

## 🎯 Learning Objectives

By the end of this lesson, I should understand:

- The basic structure of a Go program.
- The purpose of `package main`.
- How import statements work.
- The role of the `main()` function.
- Statements, expressions, and comments.
- Go formatting using `gofmt`.
- The difference between `go run` and `go build`.

---

## 1. What Is Program Structure?

Program structure describes how source code is organized so that the compiler can understand it and developers can maintain it.

A typical Go source file contains:

1. Package declaration
2. Import statements
3. Declarations such as variables, constants, and functions
4. Executable statements inside functions
5. Comments, when needed

Not every program needs all these elements.

## 2. Basic Go Program

```go
package main

import "fmt"

// main is the entry point of this program.
func main() {
    fmt.Println("Welcome to Golang!")
    fmt.Println("Day 3: Program Structure")
    fmt.Println("Learning Go step by step.")
}
```

### Output

```text
Welcome to Golang!
Day 3: Program Structure
Learning Go step by step.
```

### Explanation

- `package main` declares the package for an executable program.
- `import "fmt"` makes the `fmt` package available in this file.
- `func main()` declares the program's entry-point function.
- `fmt.Println()` prints text and adds a newline.
- `{}` delimit the function body.
- `//` begins a single-line comment.

## 3. Package Declaration

A package organizes related Go source code.

```go
package main
```

For a normal executable:

- The package must be named `main`.
- The package must provide a `main()` function.
- Execution begins in `main()` after required initialization.

Reusable packages commonly use other names, such as `strings` or `mathutil`.

## 4. Import Statements

An import makes another package available to the current source file.

```go
import "fmt"
```

The `fmt` package provides formatted input and output functions.

### Importing Multiple Packages

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Hello, Snegan!")
    fmt.Println(strings.ToUpper("learning golang"))
}
```

### Output

```text
Hello, Snegan!
LEARNING GOLANG
```

**Important:** Go reports an error when a package is imported but not used in that file.

## 5. The `main()` Function

```go
func main() {
    fmt.Println("Program started")
}
```

- `func` declares a function.
- `main` is the function name.
- `()` contains the parameter list.
- `{}` define the function body.

A Go program can contain multiple functions, but the executable starts from `main()`.

Example:

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello!")
}

func main() {
    greet()
    fmt.Println("Program finished")
}
```

### Output

```text
Hello!
Program finished
```

The `greet()` function runs because `main()` calls it.

## 6. Statements and Expressions

An **expression** produces a value.

Example:

```go
10 + 20
```

The expression evaluates to `30`.

A **statement** performs an action.

Example:

```go
fmt.Println(10 + 20)
```

This function-call statement prints `30`.

Go generally does not require developers to write semicolons at the end of each line. The language's semicolon insertion rules handle this automatically in normal Go code.

## 7. Comments in Go

Comments help developers understand code. They are not executable instructions.

### Single-Line Comment

```go
// Print a welcome message.
fmt.Println("Welcome")
```

### Multi-Line Comment

```go
/*
This is a multi-line comment.
It explains the purpose of the code.
*/
```

**Best practice:** Write comments that explain intent or non-obvious decisions.

## 8. Formatting Go Code

Go provides `gofmt` to format code consistently.

Format a file:

```bash
gofmt -w day-03/main.go
```

- `gofmt` formats Go code.
- `-w` writes the changes to the file.

Consistent formatting improves readability and makes collaboration easier.

## 9. `go run` vs `go build`

### Run a Go Program

```bash
go run ./day-03
```

This builds and runs the program conveniently.

### Build an Executable

```bash
go build -o day03 ./day-03
```

This creates an executable named `day03` in the current directory.

Run it on Ubuntu:

```bash
./day03
```

| Command | Purpose |
|---|---|
| `go run ./day-03` | Builds and runs the program |
| `go build -o day03 ./day-03` | Builds an executable |
| `./day03` | Runs the executable |
| `gofmt -w day-03/main.go` | Formats the source file |

## 10. How a Go Program Executes

A simplified execution flow is:

```text
Go Source Code
      |
      v
Go Toolchain
      |
      v
Executable Binary
      |
      v
Operating System Loads Program
      |
      v
Go Runtime Initialization
      |
      v
main() Executes
      |
      v
Program Output
```

The compiler translates Go source code into machine code as part of the build process. The operating system loads and manages the running process, while the Go runtime provides language-specific runtime services.

## 11. Common Beginner Mistakes

### Mistake 1: Incorrect Package Name

```go
package Main
```

`Main` and `main` are different identifiers. Use `package main` for an executable.

### Mistake 2: Missing Import

Calling `fmt.Println()` without importing `fmt` causes a compilation error.

### Mistake 3: Unused Import

Importing a package without using it causes a compilation error.

### Mistake 4: Missing Closing Brace

Every opening brace must have a matching closing brace.

### Mistake 5: Incorrect Command Location

Run `go run ./day-03` from the project root, where `go.mod` is located.

---

## 12. Practice Exercises

### Exercise 1 — Personal Introduction

Write a program that prints:

```text
My name is Snegan.
I am learning Golang.
My goal is to become a skilled software engineer.
```

Requirements:

- Use `package main`.
- Import `fmt`.
- Write a `main()` function.
- Use `fmt.Println()`.

### Exercise 2 — Multiple Imports

Import `fmt` and `strings`.

Print the uppercase version of:

```text
golang program structure
```

Hint: Use `strings.ToUpper()`.

### Exercise 3 — Comments

Add a single-line comment and a multi-line comment to a working Go program.

Run it and verify that comments do not appear in the output.

### Exercise 4 — Debugging

Temporarily remove an import or closing brace.

Run the program, read the error message, and fix the problem.

### Exercise 5 — Build and Execute

Run these commands from the project root:

```bash
go run ./day-03
go build -o day03 ./day-03
./day03
rm day03
```

Observe the difference between running source code conveniently and executing a separately built binary.

---

## 13. Interview Questions

1. What is a package in Go?
2. Why do we use `package main`?
3. What is the role of the `main()` function?
4. Why do we import packages?
5. What is the difference between a statement and an expression?
6. Why does Go report unused imports?
7. What is the purpose of comments?
8. What does `gofmt` do?
9. What is the difference between `go run` and `go build`?
10. Can a Go executable contain multiple functions?

### Short Answers

1. A package organizes related Go source code.
2. `package main` declares the package for an executable program.
3. `main()` is the executable's entry point.
4. Imports make other packages available to a source file.
5. An expression produces a value; a statement performs an action.
6. Go requires imported packages to be used in the file.
7. Comments explain code to developers.
8. `gofmt` formats Go source code consistently.
9. `go run` builds and runs a program; `go build` builds an executable.
10. Yes. The executable begins at `main()`, which can call other functions.

---

## 14. Day 3 Summary

Today I learned:

- The structure of a Go source file.
- Package declarations and imports.
- The `main()` function.
- Statements, expressions, and comments.
- Go formatting with `gofmt`.
- Running and building Go programs.
- The basic flow from source code to execution.

## 15. Completion Checklist

- [ ] I understand `package main`.
- [ ] I understand import statements.
- [ ] I can explain the `main()` function.
- [ ] I can write and run a Go program.
- [ ] I understand comments, statements, and expressions.
- [ ] I can use `gofmt`.
- [ ] I can build and run an executable.
- [ ] I completed the practice exercises.
- [ ] I reviewed the interview questions.

---

**Author:** Snegan  
**GitHub:** [@devsnegan](https://github.com/devsnegan)  
**Instagram:** @dev_snegan

**Next Topic:** Day 4 — Variables & Constants