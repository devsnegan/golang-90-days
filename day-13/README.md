# Day 13 - Strings and Runes in Go

## Overview

Today I learned how Go handles strings, bytes, and Unicode characters.

## Topics Covered

1. What is a string?
2. String immutability
3. String length using `len()`
4. UTF-8 encoding
5. What is a byte?
6. What is a rune?
7. Difference between `[]byte` and `[]rune`
8. Convert string to byte slice
9. Convert string to rune slice
10. Iterate over strings using `range`
11. Modify byte slices and rune slices
12. Reverse a string using runes

## String

A string in Go is an immutable sequence of bytes.

```go
text := "Golang"
fmt.Println(text)
fmt.Println(len(text))
```

## Byte Slice

A byte is an alias for `uint8`.

```go
text := "Go"
bytes := []byte(text)

fmt.Println(bytes)
```

Output:

```text
[71 111]
```

## Rune

A rune is an alias for `int32` and represents a Unicode code point.

```go
text := "க"
runes := []rune(text)

fmt.Println(runes)
```

## Byte Slice vs Rune Slice

| Feature | []byte | []rune |
|---|---|---|
| Element type | uint8 | int32 |
| Represents | Bytes | Unicode code points |
| Element size | 1 byte | 4 bytes |
| Common use | Files, HTTP, binary data | Unicode text processing |
| len() | Byte count | Rune count |

## Important Difference

```go
text := "Goக😊"

fmt.Println(len(text))
fmt.Println(len([]rune(text)))
```

Output:

```text
9
4
```

`len(string)` returns the number of bytes, while `len([]rune(string))` returns the number of Unicode code points.

## Practical Applications

- Reading HTTP request bodies
- Processing JSON data
- Working with files
- Handling multilingual user names
- Processing Unicode text
- Counting Unicode code points
- Reversing text by code point

## Exercises Completed

- [x] Exercise 01 - Basic String
- [x] Exercise 02 - Byte Slice Conversion
- [x] Exercise 03 - Rune Slice Conversion
- [x] Exercise 04 - Count Bytes and Runes
- [x] Exercise 05 - Iterate Over String
- [x] Exercise 06 - Reverse String Using Runes

## Key Takeaways

- Strings are immutable.
- Go strings use UTF-8 encoding for text.
- `len(string)` returns bytes.
- `[]byte` represents the bytes of a string.
- `[]rune` represents Unicode code points.
- `range` over a string decodes UTF-8 into runes.

## Next Day

Day 14 - Functions in Go