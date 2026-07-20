# Requirements: `calc` package with an `Add` function

## Scope

Introduce a self-contained `calc` package (import path
`github.com/hanahmily/tdd-sandbox/calc`) exposing a single public operation:
integer addition. The contract — public interface plus RED unit and
end-to-end tests — is defined first; production behavior is filled in by a
later `implement` step.

## Functional requirements

- **R1 — Signature.** The package MUST export `Add(a, b int) int`.
- **R2 — Definition.** `Add(a, b)` MUST return the ordinary integer sum,
  `a + b`, for the full range of `int` inputs.
- **R3 — Value classes.** `Add` MUST be correct for positive operands,
  negative operands, zero operands, and mixed-sign operands.
- **R4 — Commutativity.** Addition MUST be commutative:
  `Add(a, b) == Add(b, a)` for every pair of operands.
- **R5 — Zero identity.** Zero MUST be the additive identity element:
  `Add(n, 0) == n` for every `n`.

## Contract & test requirements

- **R6 — Approved interface.** `calc/calc.go` declares the approved public
  interface — the `Add(a, b int) int` signature and its documented contract —
  as a stub with **no production behavior** (it `panic`s). The signature and
  doc comment are FINAL; the `implement` step replaces ONLY the function body
  with `return a + b`. Because Go permits exactly one declaration of `Add` and
  it must live in this file, `calc/calc.go` is deliberately **NOT** in the
  protected set — the required body swap has to happen here. The approved
  signature and semantics are instead pinned (and thereby protected) by the RED
  tests in R7/R8, which fail to compile or fail to pass if the interface drifts.
- **R7 — RED unit tests.** `calc/calc_test.go` (internal `calc` package) pins
  the value classes (R3), commutativity (R4), and zero identity (R5) directly
  against the package. These tests are RED against the panic stub.
- **R8 — RED end-to-end test.** `calc/e2e_test.go` (external `calc_test`
  package) exercises `Add` as a real consumer would — folding a stream of
  signed transactions into a running ledger balance through the public API. The
  e2e test is RED against the stub, and its name (`TestE2ELedger`) matches the
  `TestE2E` gate pattern.
- **R9 — Protected artifacts.** The unit-test (`calc/calc_test.go`) and e2e
  (`calc/e2e_test.go`) files are listed in `.vajra-protected`; the `implement`
  step MUST NOT modify them. `calc/calc.go` is intentionally left unprotected:
  the implement step replaces its stub body (and ONLY the body) with
  `return a + b`, leaving the exported signature and doc comment untouched. Any
  signature drift breaks compilation of the protected tests and is caught
  immediately.

## Red/green gates

- Assign (now): `go test ./...` and `go test -run TestE2E ./...` MUST fail —
  the panic stub keeps every assertion RED.
- Implement (later): `go test ./...` and `go test -cover ./...` MUST pass once
  the body becomes `return a + b`.

## Out of scope

- Any arithmetic operation other than addition.
- Overflow handling beyond native Go `int` semantics.
- Production implementation of `Add` (delivered by the `implement` step, which
  replaces the stub body of the unprotected `calc/calc.go` with `return a + b`
  and nothing more — signature and doc comment stay exactly as declared).
