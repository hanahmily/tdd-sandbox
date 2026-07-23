# Package `difference` — Interface Contract

`difference` provides integer subtraction for the tdd-sandbox milestone.

## Public API

```go
package difference

// Sub returns the arithmetic difference a - b.
//
// It operates on the full signed int range and performs no overflow
// checking (matching Go's native integer subtraction semantics). The
// result is exactly a - b for every pair of operands.
func Sub(a, b int) int
```

## Behavioral contract

| Case                         | Example              | Expected     |
|------------------------------|----------------------|--------------|
| Positive operands            | `Sub(9, 4)`          | `5`          |
| Result is negative           | `Sub(4, 9)`          | `-5`         |
| Negative operands            | `Sub(-3, -8)`        | `5`          |
| Mixed signs                  | `Sub(-3, 5)`         | `-8`         |
| Zero subtrahend (identity)   | `Sub(7, 0)`          | `7`          |
| Zero minuend                 | `Sub(0, 7)`          | `-7`         |
| Both zero                    | `Sub(0, 0)`          | `0`          |
| Equal operands               | `Sub(6, 6)`          | `0`          |

## Notes

- `Sub` is a pure function: no state, no side effects, deterministic.
- Overflow follows Go's wraparound semantics and is out of scope for
  this milestone (no operands near `math.MaxInt` are exercised).
