// Package calc provides basic integer arithmetic helpers.
//
// The milestone delivers a single exported operation, Add, whose contract is
// fixed by this file. Add must return the ordinary signed-integer sum of its
// two arguments and must handle negative and zero operands correctly.
package calc

// Add returns the sum of its two integer arguments.
//
// For every pair of ints a and b the result is the ordinary two's-complement
// signed sum a + b. In particular:
//   - Add(x, 0) == x and Add(0, x) == x (zero is the additive identity).
//   - Negative operands are summed like any other value, so
//     Add(-4, -6) == -10 and Add(-5, 5) == 0.
//   - Add is commutative: Add(a, b) == Add(b, a).
//
// This file is the approved public interface: the exported signature
//
//	func Add(a, b int) int
//
// is fixed and must not change. Behaviour is supplied by the unexported add
// helper, which the implementer must provide in a separate, non-protected file
// (for example calc/add.go):
//
//	package calc
//
//	func add(a, b int) int { return a + b }
func Add(a, b int) int {
	return add(a, b)
}
