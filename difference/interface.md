# `difference` package — interface contract

This document is the approved public interface for the `difference` milestone.
It is protected: the coder implements against it and must not alter the
signature or documented semantics.

## Public API

```go
package difference

// Sub returns the difference a - b.
func Sub(a, b int) int
```

## Semantics

`Sub(a, b)` returns the integer difference `a - b`, using Go's native `int`
arithmetic. There is no error path and no overflow handling beyond Go's
defined two's-complement wraparound for `int`.

The contract must hold across all three operand classes:

- **Positive operands** — e.g. `Sub(5, 3) == 2`; `Sub(3, 5) == -2`.
- **Negative operands** — e.g. `Sub(-5, -3) == -2`; `Sub(-4, 6) == -10`.
- **Zero operands** — e.g. `Sub(0, 0) == 0`; `Sub(0, 7) == -7`; `Sub(7, 0) == 7`.

## Test surface

- Unit: `difference/difference_test.go` → `TestSub` (table-driven, white-box).
- End-to-end: `e2e/e2e_test.go` → `TestE2EDifferenceSub` (imports the package
  by module path as an external consumer).
