# learn-go

A place to familiarize with Go syntax and features.

## Packages

### `basics`
Covers core Go syntax:
- Variables, constants, and basic types
- Control flow (`if`/`else`, `switch`)
- Loops and range
- Functions: multiple return values, variadic, named returns
- Closures
- Slices, maps
- Structs and interfaces
- Pointers
- Error handling with custom error types
- Defer

### `concurrency`
Covers Go concurrency primitives:
- Goroutines and channels
- WaitGroups and Mutexes
- Fan-out pattern
- Pipeline pattern

## Running tests

```sh
go test ./...
```