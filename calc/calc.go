// Package calc provides small, well-scoped integer arithmetic helpers.
package calc

// Mul returns the product of a and b.
//
// This is the approved public interface for the multiplication utility. The
// signature is frozen: the production behavior is supplied by the unexported
// mul function (see mul.go), which the implementer completes. Do not change
// this file — it defines the contract that the unit and e2e tests verify.
func Mul(a, b int) int {
	return mul(a, b)
}
