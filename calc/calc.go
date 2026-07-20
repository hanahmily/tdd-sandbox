// Package calc provides basic integer arithmetic helpers.
//
// This milestone delivers a single addition primitive that is correct for the
// full range of int inputs, including negative operands and zero.
package calc

// Add returns the sum of a and b.
//
// It is defined over all int values: the two operands may be positive,
// negative, or zero, in any combination. Add is a pure function — it has no
// side effects and always returns a+b.
func Add(a, b int) int {
	return a + b
}
