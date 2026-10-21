# Day 6 — Type Conversion in Go

**90 Days of Golang Mastery**  
**Author:** dev_snegan

## Overview

Type conversion changes a value from one data type to another. Go requires explicit conversions between different numeric types.

## Learning Objectives

- Understand explicit type conversion.
- Convert between integer and floating-point types.
- Understand truncation when converting floats to integers.
- Convert numbers to strings using `strconv`.
- Parse strings into numbers.
- Handle invalid numeric input.
- Apply conversion to a backend pagination example.

## 1. Numeric Type Conversion

```go
var age int = 25
ageAsFloat := float64(age)
```

Go does not automatically convert incompatible numeric types for arithmetic.

## 2. Float to Integer

```go
price := 99.99
converted := int(price)

fmt.Println(converted) // 99
```

Conversion truncates the fractional part toward zero rather than rounding.

## 3. Integer to String

Use `strconv.Itoa` to convert an integer to its decimal string representation.

```go
import "strconv"

number := 2026
text := strconv.Itoa(number)
```

Avoid using `string(number)` when the goal is to produce decimal digits. That conversion interprets the integer as a Unicode code point.

## 4. String to Integer

Use `strconv.Atoi` and check the returned error.

```go
number, err := strconv.Atoi("150")
if err != nil {
    // Handle invalid input.
    return
}

fmt.Println(number)
```

## 5. String to Float

Use `strconv.ParseFloat`.

```go
price, err := strconv.ParseFloat("149.50", 64)
if err != nil {
    // Handle invalid input.
    return
}

fmt.Println(price)
```

## 6. Practical Backend Example

A pagination API may receive its page number as a string. Parse it into an integer before calculating the offset.

```go
pageText := "4"
page, err := strconv.Atoi(pageText)
if err != nil || page < 1 {
    // Reject invalid input.
    return
}

pageSize := 20
offset := (page - 1) * pageSize
```

In production, also validate upper limits and check arithmetic for overflow when input values can be large.

## Exercises Completed

1. Integer to float conversion.
2. Float to integer conversion.
3. Integer to string conversion.
4. String to integer conversion.
5. Fixing a numeric type mismatch.
6. Parsing a pagination parameter and calculating an offset.

## Interview Questions

1. What is type conversion in Go?
2. Why does Go require explicit numeric conversions?
3. What happens when a float is converted to an integer?
4. What is the difference between `string(65)` and `strconv.Itoa(65)`?
5. How does `strconv.Atoi` handle invalid input?
6. How do you parse a string into a `float64`?
7. Why should numeric ranges be validated before conversion?

## Key Takeaways

- Numeric conversions are explicit.
- Float-to-integer conversion truncates toward zero.
- Narrowing integer conversions can lose information.
- `strconv` provides useful numeric parsing and formatting functions.
- Always handle errors when parsing external input.

## Completion Checklist

- [ ] Main program completed.
- [ ] Six exercises completed.
- [ ] README added.
- [ ] Programs formatted and executed.
- [ ] Git changes reviewed and pushed.

**Next:** Day 7 — Operators

---

*Learning Go step by step — @dev_snegan*
