# Requirements — `calc.Mul` multiplication utility

Milestone: add a tiny, well-scoped multiplication utility to the `calc`
package. The contract is intentionally minimal: one exported function plus a
real end-to-end use case that consumes it.

Each requirement below is numbered and mapped to the test(s) that prove it.

| #   | Requirement | Proven by |
| --- | ----------- | --------- |
| R1  | The `calc` package exposes an exported function with the exact signature `func Mul(a, b int) int`. | Compilation of `calc/calc.go` (interface) + all tests referencing `calc.Mul` |
| R2  | `Mul` returns the arithmetic product of two positive integers. | `TestMul/two_positives` (`calc/calc_test.go`) |
| R3  | `Mul` returns `0` when either operand is `0`. | `TestMul/multiply_by_zero` (`calc/calc_test.go`) |
| R4  | `Mul` returns the other operand unchanged when one operand is `1` (multiplicative identity). | `TestMul/positive_times_one` (`calc/calc_test.go`) |
| R5  | `Mul` handles a negative operand, producing a negative product. | `TestMul/negative_times_positive` (`calc/calc_test.go`) |
| R6  | `Mul` handles two negative operands, producing a positive product. | `TestMul/two_negatives` (`calc/calc_test.go`) |
| R7  | `Mul` is commutative: `Mul(a, b) == Mul(b, a)`. | `TestMulIsCommutative` (`calc/calc_test.go`) |
| R8  | A real consumer can import `calc` across a package boundary and use `Mul` to compute a cart order total (sum of quantity × unit price per line). | `TestE2EOrderTotal` (`e2e/e2e_test.go`) |

## Notes

- The public interface (`calc/calc.go`) is frozen and protected. It forwards to
  an unexported `mul` in `calc/mul.go`, which is the only production file the
  implementer edits.
- All listed tests are RED at contract time (the `mul` stub panics) and must go
  GREEN once `mul` is implemented as integer multiplication.
- The e2e gate runs `go test -run TestE2E ./...`; `TestE2EOrderTotal` matches
  that prefix.
