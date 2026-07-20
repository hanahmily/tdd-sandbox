# `calc` package milestone contract

## Purpose

Provide a small, reusable integer-addition operation for callers that need to combine two values. The milestone is intentionally limited to the package-level `Add` function; it does not introduce state, configuration, or a broader calculator API.

## Requirements

1. The module MUST expose a package at `github.com/hanahmily/tdd-sandbox/calc`.
2. The package MUST export a function with the exact signature `func Add(a, b int) int`.
3. `Add` MUST return the sum of its two arguments using Go's `int` arithmetic semantics.
4. `Add` MUST handle positive, negative, mixed-sign, and zero operands correctly. Zero MUST behave as the additive identity.
5. For the same pair of inputs, `Add` MUST return the same result and MUST NOT require caller-managed state or configuration.
6. The API MUST return an `int` directly and MUST NOT return an error for ordinary integer addition.
7. The milestone MUST remain dependency-free beyond the Go standard toolchain and MUST NOT add unrelated calculator operations or observable side effects.

## Approved public interface

The only approved public behavior is the interface documented in [`calc/interface.md`](calc/interface.md). Names, parameter types, return type, and package path are part of the contract.

## Acceptance criteria

- Unit coverage exercises positive values, negative values, mixed signs, and zero on either side of the operation.
- An end-to-end scenario composes the operation in a realistic account-balance calculation involving a deposit and a negative withdrawal.
- The tests are committed before production behavior exists, so the initial test run is expected to be RED because `Add` is not implemented yet.
- No production implementation is included in this milestone's assignment.
