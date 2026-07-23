# Interface contract: package `difference`

Import path: `github.com/hanahmily/tdd-sandbox/difference`

## Public API

```go
// Sub returns the difference of a and b (a - b).
func Sub(a, b int) int
```

## Semantics

- `Sub(a, b)` computes the integer `a - b`.
- Pure function: no state, no side effects, deterministic.
- Must be correct for positive operands, negative operands (including a
  negative result), and zero operands.

## Frozen contract

The exported signature `func Sub(a, b int) int` and the import path are the
approved contract and MUST NOT change during implementation. The function
*body* in `difference/difference.go` is mutable and is what the coder fills
in to turn the suite green.
