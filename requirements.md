# Requirements — `calc` package with `Add`

## Summary

Deliver a Go package `calc` exposing a single exported operation, `Add`, that
returns the signed-integer sum of its two arguments and handles negative and
zero operands correctly.

## Approved public interface (frozen)

The public surface is fixed by `calc/calc.go` and MUST NOT change:

```go
func Add(a, b int) int
```

`Add` delegates to an unexported helper `add(a, b int) int` that the
implementer supplies in a separate, non-protected file (for example
`calc/add.go`). The exported signature and its documented guarantees are the
approved contract; only the helper is left to implement.

## Functional requirements

For all pairs of `int` values `a` and `b`, `Add(a, b)` returns the ordinary
two's-complement signed sum `a + b`. This single rule yields:

- **R1 — Sum:** `Add(a, b) == a + b`.
- **R2 — Zero identity:** zero is the additive identity from both sides, i.e.
  `Add(x, 0) == x` and `Add(0, x) == x`.
- **R3 — Negatives:** negative operands are summed like any other value, e.g.
  `Add(-4, -6) == -10` and `Add(-5, 5) == 0`.
- **R4 — Mixed signs:** mixed-sign operands sum correctly, e.g.
  `Add(-3, 8) == 5` and `Add(10, -25) == -15`.
- **R5 — Commutativity:** `Add(a, b) == Add(b, a)`.
- **R6 — Large operands:** ordinary (non-overflowing) large magnitudes sum
  correctly, e.g. `Add(1_000_000, 2_345_678) == 3_345_678`.

## RED test suites (frozen)

The following files encode the contract and are RED at handoff — the package
does not yet compile because the `add` helper is absent. Turning them green is
the definition of done for the implementation step. They are listed in
`.vajra-protected` and MUST NOT be modified by the implementer.

- **Unit tests — `calc/calc_test.go`** (external `calc_test` package, public API
  only): `TestAdd` covers positive, negative, zero, mixed-sign and large
  operands (R1, R3, R4, R6); `TestAddZeroIdentity` covers R2; `TestAddCommutative`
  covers R5.
- **End-to-end test — `calc/e2e_test.go`**: `TestLedgerBalanceE2E` folds
  `calc.Add` over a stream of deposits (positive), withdrawals (negative) and
  no-op (zero) transactions to compute a running balance, exercising the whole
  contract the way a real consumer would; `TestReconciliationE2E` replays the
  same stream from both ends and asserts the balance is order-independent
  (a real reconciliation use case that also reinforces R5).

## Out of scope / non-goals

- No overflow handling beyond Go's defined two's-complement wraparound; inputs
  are assumed within `int` range.
- No additional operations, error returns, or exported symbols beyond `Add`.

## Definition of done

The implementer supplies the `add` helper in a non-protected file; both the
unit and end-to-end suites pass with the approved public interface and the
protected test files left unchanged.
