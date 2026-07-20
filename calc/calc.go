// Package calc provides basic integer arithmetic operations.
package calc

// Add returns the sum of a and b.
//
// The result is the ordinary integer sum, so negative and zero operands are
// handled correctly: Add(-2, 3) == 1, Add(0, 0) == 0, Add(-4, -6) == -10 and
// Add(8, -8) == 0.
//
// Add is the frozen public contract for this milestone. Its behavior is
// provided by the unexported add function, which the implementation must
// complete.
func Add(a, b int) int {
	return add(a, b)
}
