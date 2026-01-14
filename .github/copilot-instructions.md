# GitHub Copilot Instructions for unmango/go

## Repository Overview

This is a Go utilities library providing functional programming patterns and abstractions. The repository is a collection of experimental and useful packages for Go, including:

- **iter**: Extensions to Go's standard `iter` package with functional operations (Map, Filter, Fold, etc.)
- **result**: Result type representing success or error
- **rx**: Reactive programming with observables and signals
- **option/maybe**: Optional value types
- **maps/slices**: Convenience wrappers for standard library packages
- **codec**: Encoding/decoding abstractions
- **os**: Operating system abstractions

## Build and Test

### Building
```bash
go build ./...
# or
make build
```

### Testing
```bash
# Run tests (excludes E2E by default locally)
make test

# Run all tests including E2E
make test_all

# Run tests with Ginkgo directly
go tool ginkgo run -r ./
```

### Other Commands
```bash
# Tidy dependencies
make tidy
go mod tidy

# Clean test artifacts
make clean
```

## Code Style and Conventions

### General Guidelines
- Use **tabs for indentation** (except YAML/Nix files which use spaces)
- Insert final newline in all files
- Trim trailing whitespace
- Follow standard Go conventions and idioms
- Keep code simple and readable

### Package Structure
- Each package should have a `_suite_test.go` file for Ginkgo test setup
- Test files use `_test.go` suffix
- Use descriptive names for functions following Go conventions (e.g., `Map`, `Filter`, `Bind`)

### Functional Programming Patterns
- Prefer immutable operations
- Functions should be pure where possible
- Use generic types with type parameters (e.g., `[V any]`, `[T, V any]`)
- Sequences (`Seq[V]`) use closure-based iterators with `yield` functions
- Return zero values with errors when operations fail

### Testing
- Use **Ginkgo v2** testing framework with **Gomega** matchers
- Test files use dot imports for Ginkgo/Gomega:
  ```go
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
  ```
- Structure tests with `Describe` and `It` blocks
- Use descriptive test names (e.g., "should append a value", "should bind")
- Use Gomega matchers like `Expect().To()`, `ConsistOf()`, `HaveExactElements()`

### Error Handling
- Return errors explicitly using `(T, error)` pattern
- Use `errors.New()` for simple errors
- Use `fmt.Errorf()` for formatted errors
- In result types, wrap operations in functions: `func() (T, error)`

### Type Aliases and Generics
- Use type aliases to re-export standard library types when needed
- Generic functions should use clear, single-letter type parameters (V, T, X, A)
- When adapting standard library, maintain compatibility while extending functionality

## Dependencies

- Go version is specified in `go.mod` (currently 1.25.5)
- Uses `gomod2nix` for Nix integration
- Testing: Ginkgo v2, Gomega
- Use `go mod tidy` to manage dependencies

## CI/CD

- CI runs on GitHub Actions (`.github/workflows/ci.yml`)
- CI includes: build, test, Nix checks, and gomod2nix validation
- Coverage reports are uploaded to Codecov
- Use `--race` flag and generate coverage profile in CI

## Nix Integration

- The project uses Nix flakes for reproducible builds
- `flake.nix` and `gomod2nix.toml` define the Nix environment
- Run `nix build` to build with Nix
- Run `nix flake check --all-systems` to verify Nix configuration
