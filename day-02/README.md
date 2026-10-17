# Day 2/90 — Go Installation & Development Setup

## 🎯 Goal

Understand the Go development environment, verify the Go installation, learn essential Go commands, and prepare Ubuntu and VS Code for Go development.

## 1. What Is a Development Environment?

A development environment is a collection of tools used to write, build, run, test, and maintain software.

Our Go development environment includes:

- **Go:** Programming language and development toolchain.
- **VS Code:** Code editor.
- **Ubuntu:** Operating system.
- **Terminal:** Interface for executing commands.
- **Git:** Version control system.
- **GitHub:** Platform for hosting our repository.

## 2. Verify Go Installation

Check the installed Go version:

```bash
go version
```

Example output:

```text
go version go1.24.0 linux/amd64
```

The version above is only an example. Your installed version may differ.

Check which Go executable is being used:

```bash
which go
```

This displays the path to the Go executable found by the shell.

## 3. Essential Go Environment Variables

### GOROOT

`GOROOT` identifies the Go distribution directory used by the toolchain.

Check it:

```bash
go env GOROOT
```

Example:

```text
/usr/local/go
```

### GOPATH

`GOPATH` identifies a workspace location used for certain Go development files, downloaded modules, and installed tools.

Check it:

```bash
go env GOPATH
```

Example:

```text
/home/snegan/go
```

Modern Go projects use modules and do not need to be located inside `GOPATH`.

### PATH

`PATH` is a shell environment variable containing directories searched when we execute commands.

Check it:

```bash
echo "$PATH"
```

If the Go executable's directory is missing from `PATH`, the terminal may not recognize the `go` command.

### Comparison

| Variable | Purpose |
|---|---|
| `GOROOT` | Go distribution directory |
| `GOPATH` | Workspace and certain Go development files |
| `PATH` | Directories searched for executable commands |

## 4. Useful Go Commands

| Command | Purpose |
|---|---|
| `go version` | Display the installed Go version |
| `go env` | Display Go environment settings |
| `go run ./day-01` | Build and run the Day 1 package |
| `go build ./day-01` | Build the Day 1 package |
| `go fmt ./...` | Format Go packages in the module |
| `go vet ./...` | Check for suspicious code patterns |
| `go test ./...` | Run tests across packages |
| `go help` | Display Go command help |
| `go mod tidy` | Reconcile module requirements and checksums |

## 5. Understanding go run and go build

### Using go run

```bash
go run ./day-01
```

This builds and runs the specified Go package, making it convenient during development.

### Using go build

```bash
go build -o day01 ./day-01
```

This creates an executable named `day01` on a typical Ubuntu setup.

Run the executable:

```bash
./day01
```

Remove the generated executable when finished:

```bash
rm day01
```

**Key difference:** `go run` builds and runs the program as a development convenience, while `go build` builds the program so it can be executed separately.

## 6. Understanding go.mod

A Go module is a collection of related Go packages managed as a unit.

The `go.mod` file declares the module path and records the Go version and dependency requirements.

Example:

```go
module github.com/devsnegan/golang-90-days

go 1.xx
```

The Go version is illustrative. Keep the version already present in your project.

Inspect the file:

```bash
cat go.mod
```

We do not need to run `go mod init` again because this repository already has a Go module.

## 7. Setting Up VS Code

1. Open the `golang-90-days` project.
2. Install the official Go extension published by the Go team.
3. Open a `.go` file.
4. Install any recommended Go tools if prompted.
5. Run Go programs through the integrated terminal.

Official extension: https://marketplace.visualstudio.com/items?itemName=golang.Go

## 8. Practical Experiments

Run these commands from the repository root:

```bash
cd ~/golang-90-days

go version
go env GOOS GOARCH GOROOT GOPATH
go env GOMOD
which go

go run ./day-01
go build -o day01 ./day-01
./day01
rm day01

go fmt ./...
go vet ./...
go test ./...
```

These experiments help verify the installation, inspect the environment, build and run the program, format code, and perform basic checks.

Note: `go test ./...` may report that packages have no test files. That is normal if tests have not been written yet.

## 9. Common Problems

| Problem | Possible solution |
|---|---|
| `go: command not found` | Check the Go installation and `PATH`. |
| `go.mod file not found` | Navigate to the repository root. |
| Wrong Go version | Check `which go`, `go version`, and `go env GOROOT`. |
| `code: command not found` | Open VS Code from the applications menu or configure its command-line launcher. |
| Multiple `main` declarations | Put independent executable exercises in separate directories. |

Always read the error message before changing your configuration.

## 10. Interview Questions

**Q1. What is a Go toolchain?**

A collection of tools used to build, run, format, test, and inspect Go programs.

**Q2. What is GOROOT?**

It identifies the Go distribution directory used by the toolchain.

**Q3. What is GOPATH?**

It identifies a workspace location used for certain Go development files, downloaded module source code, and installed tools.

**Q4. What is PATH?**

An environment variable that tells the shell where to search for executable commands.

**Q5. What is go.mod?**

A file that declares a Go module and records module-level configuration and dependency requirements.

**Q6. What is the difference between go run and go build?**

`go run` builds and runs a package as a development convenience. `go build` builds the package into an executable when building a `main` package.

**Q7. What is gofmt?**

A standard Go formatter that applies consistent formatting rules to source code.

**Q8. Why do we use go test?**

It runs tests defined in Go test files to check whether code behaves as expected.

## 11. Day 2 Summary

Today, we learned:

- What a Go development environment contains.
- How to verify the Go installation.
- The purposes of `GOROOT`, `GOPATH`, and `PATH`.
- How to inspect Go environment settings.
- The difference between `go run` and `go build`.
- The purpose of `go.mod`.
- How to use `gofmt`, `go vet`, and `go test`.
- How to prepare VS Code for Go development.

## 12. Practice Checklist

- [ ] Check the installed Go version.
- [ ] Inspect Go environment variables.
- [ ] Find the Go executable path.
- [ ] Inspect `go.mod`.
- [ ] Run the Day 1 program.
- [ ] Build and execute a binary.
- [ ] Run formatting, vetting, and test commands.
- [ ] Commit the Day 2 notes to GitHub.

---

**Learning Journey:** Day 2 of 90 Days of Golang Mastery

**Author:** Snegan

**GitHub:** @devsnegan

**Instagram:** @dev_snegan
