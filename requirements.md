# Requirements: `calc.Add` integer addition

## Goal

Add a package `calc` exposing a single arithmetic primitive that returns the
sum of its two integer arguments, handling negative and zero values correctly.

## Public contract (frozen, protected)

```go
// Package calc provides basic integer arithmetic operations.
package calc

// Add returns the sum of a and b.
func Add(a, b int) int
```

- `Add` is the frozen, protected public API (`calc/calc.go`). It delegates to
  an unexported `add` function in the non-protected `calc/calc_impl.go`, which
  is the intended place for the implementation phase to write production
  behavior.
- The published signature and documented semantics must not drift.

## Functional requirements

1. `Add(a, b)` returns the ordinary integer sum `a + b`.
2. Correct across the full sign space:
   - two positives — `Add(2, 3) == 5`
   - positive/zero and zero/positive — `Add(5, 0) == 5`, `Add(0, 7) == 7`
   - both zero — `Add(0, 0) == 0`
   - two negatives — `Add(-4, -6) == -10`
   - mixed signs — `Add(-3, 10) == 7`, `Add(-10, 3) == -7`, `Add(8, -8) == 0`
   - negative/zero — `Add(-5, 0) == -5`
   - large operands — `Add(1_000_000, 2_000_000) == 3_000_000`
3. `Add` is commutative: `Add(a, b) == Add(b, a)`.

## RED tests (frozen, protected — must pass only after implementation)

- **Unit** (`calc/calc_test.go`): table-driven cases across positive, negative,
  and zero operands, plus a commutativity check.
- **End-to-end** (`e2e/calc_e2e_test.go`): simulates a real consumer computing
  a running account balance over a ledger of deposits and withdrawals using
  only `calc.Add`, asserting the per-step and final balances, and a
  nets-to-zero drain. E2E test names are prefixed `TestE2E` so the e2e gate
  (`go test -run TestE2E ./...`) selects them.

## Out of scope

- Any operation other than integer addition.
- Overflow handling beyond native `int` semantics.

## Definition of done

- Contract, RED unit tests, and RED e2e test are frozen up front.
- The implementation phase makes both suites green without modifying any
  protected file (see `.vajra-protected`).
