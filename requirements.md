# Requirements: `calc` package with a correct `Add`

## Scope

Introduce a small, self-contained `calc` package exposing a single public
primitive, `Add(a, b int) int`, that returns the arithmetic sum of its two
operands over the **entire** `int` domain.

## Approved public interface

```go
// Package: github.com/hanahmily/tdd-sandbox/calc
func Add(a, b int) int
```

This signature is final. The implementer must not change it.

## Functional requirements

- **R1 — Sum:** `Add(a, b)` returns `a + b` for every pair of `int` values.
- **R2 — Positives:** two positive operands sum correctly (e.g. `2 + 3 == 5`).
- **R3 — Negatives:** two negative operands sum correctly (e.g. `-4 + -6 == -10`).
- **R4 — Mixed signs:** operands of opposite sign sum correctly, including exact
  cancellation (e.g. `8 + (-8) == 0`, `3 + (-10) == -7`).
- **R5 — Zero identity:** zero is the additive identity for both positive and
  negative operands (`Add(n, 0) == n` and `Add(0, n) == n`).
- **R6 — Commutativity:** operand order is irrelevant — `Add(a, b) == Add(b, a)`.

## Non-functional requirements

- **N1 — Purity:** `Add` is pure and total — no side effects, defined for every
  `int` pair, no special-casing, no error path.

## Test strategy (RED)

The contract ships with tests written first and intentionally failing against an
unimplemented stub (`calc.Add` panics):

- **Unit tests** (`calc/calc_test.go`) drive `Add` directly with a table
  covering R2–R5 plus commutativity (R6) and the additive identity (R5).
- **End-to-end test** (`e2e/e2e_test.go`, `TestE2E*`) imports the package
  through its public module path and uses `Add` as a downstream consumer would:
  folding a ledger of deposits, withdrawals, and no-op adjustments into a
  running balance that passes through zero and into the negative.

The implementer's job is to replace the stub body so both suites turn green,
without modifying the approved interface, the unit tests, or the e2e test.
