# Day 25 - Type Switches in Go

## What is a Type Switch?

A type switch checks the dynamic type stored inside
an interface variable.

## Syntax

```go
switch v := value.(type) {
case string:
    // Handle string
case int:
    // Handle integer
default:
    // Handle other types
}