# Day 4 — Variables & Constants in Go

**Phase:** 1 — Go Fundamentals  
**Day:** 4 of 90  
**Topic:** Variables & Constants

## Learning Objectives

By the end of this lesson, I should understand:

- What variables are and why they are used.
- Variable declarations using `var` and `:=`.
- Variable assignment and reassignment.
- Zero values in Go.
- Constants using `const`.
- Typed and untyped constants.
- Multiple variable declarations.
- Naming rules and common mistakes.

## 1. What Is a Variable?

A variable is a named value that a program can access and, when permitted, change.

```go
var name string = "Snegan"
var age int = 25
```

Here:

- `var` declares a variable.
- `name` and `age` are identifiers.
- `string` and `int` are types.
- `"Snegan"` and `25` are initial values.

## 2. Variable Declaration Methods

### Explicit Type

```go
var name string = "Snegan"
var age int = 25
```

### Type Inference

```go
var name = "Snegan"
var age = 25
```

Go infers the types from the initial values.

### Short Variable Declaration

```go
name := "Snegan"
age := 25
```

The `:=` operator declares and initializes variables inside functions. It cannot be used directly at package level.

## 3. `var` vs `:=`

| Feature | `var` | `:=` |
|---|---|---|
| Explicit type | Supported | Not specified explicitly |
| Inside functions | Yes | Yes |
| At package level | Yes | No |
| Declaration without an initial value | Yes | No |

Example:

```go
var score int
score = 100
```

Compared with:

```go
score := 100
```

## 4. Zero Values

Go initializes variables declared without initial values using their type's zero value.

| Type | Zero value |
|---|---|
| `int` | `0` |
| `float64` | `0` |
| `string` | Empty string |
| `bool` | `false` |
| Pointer | `nil` |
| Slice | `nil` |
| Map | `nil` |
| Interface | `nil` |

Example:

```go
var age int
var name string
var active bool

fmt.Println(age)    // 0
fmt.Println(name)   // Empty string
fmt.Println(active) // false
```

## 5. Assignment and Reassignment

```go
balance := 1000
balance = 1500
```

- `:=` declares and initializes the variable.
- `=` assigns a new value to an existing variable.

A variable's type does not change when its value is reassigned.

## 6. Multiple Variable Declarations

```go
var (
    name   string = "Snegan"
    role   string = "Developer"
    active bool   = true
)
```

Short declaration can also initialize multiple variables:

```go
name, role := "Snegan", "Developer"
```

## 7. Constants

A constant represents a value that cannot be reassigned.

```go
const appName = "User Service"
const maxRetries = 3
```

Constants are useful for fixed values, such as application labels and mathematical constants.

This is invalid:

```go
const maxRetries = 3
maxRetries = 5
```

Use a variable if the value needs to change.

## 8. Typed and Untyped Constants

### Untyped Constant

```go
const number = 10
```

Go does not explicitly assign a type in this declaration. The constant can be used where its value is representable in the required type.

### Typed Constant

```go
const timeoutSeconds int = 30
```

The constant explicitly has type `int`.

Untyped constants provide flexibility when working with different numeric types, but they still follow Go's type and representability rules.

## 9. Naming Rules

Examples of common Go naming conventions:

```go
userName := "Snegan"
maxRetryCount := 3
```

Go identifiers are case-sensitive:

```go
name := "Snegan"
Name := "Developer"
```

These are different identifiers.

Names starting with an uppercase letter are exported when declared at package scope. Local variables inside a function are not exported to other packages.

Unused local variables cause compilation errors.

## 10. Practical Program

```go
package main

import "fmt"

func main() {
    var name string = "Snegan"
    age := 25
    const language = "Golang"

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Learning:", language)
}
```

## 11. Practice Exercises

1. Create an employee profile using variables.
2. Declare a balance and reassign it.
3. Create constants for an application name and retry limit.
4. Print the zero values of `int`, `string`, `bool`, and `float64`.
5. Fix a program that tries to reassign a constant.

Solutions are in the Day 4 exercise folders.

## 12. Interview Questions

1. What is a variable in Go?
2. What is the difference between `var` and `:=`?
3. What are zero values?
4. Can a variable change its type after declaration?
5. What is a constant?
6. What is the difference between variables and constants?
7. What is the difference between `=` and `:=`?
8. What are typed and untyped constants?
9. Can multiple variables be declared together?
10. Why does Go report unused local variables?

## 13. Day 4 Summary

Today I learned:

- Variable declaration and initialization.
- `var` and `:=`.
- Assignment and reassignment.
- Zero values.
- Constants and constant groups.
- Typed and untyped constants.
- Identifier naming rules.

## Completion Checklist

- [ ] I understand variable declarations.
- [ ] I can use `var` and `:=`.
- [ ] I understand zero values.
- [ ] I can reassign variables.
- [ ] I understand constants.
- [ ] I can explain typed and untyped constants.
- [ ] I completed all five exercises.
- [ ] I reviewed the interview questions.

**Author:** Snegan  
**GitHub:** [@devsnegan](https://github.com/devsnegan)  
**Instagram:** @dev_snegan

**Next Topic:** Day 5 — Data Types