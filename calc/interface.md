# Approved `calc` public interface

The package exposes exactly one public operation:

```go
package calc

// Add returns the sum of a and b using Go's int arithmetic semantics.
func Add(a, b int) int
```

Callers may pass any `int`, including zero and negative values. No error or configuration is part of this interface.
