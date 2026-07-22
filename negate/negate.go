// Package negate provides arithmetic negation of integers.
package negate

// Value returns the arithmetic negation of a (that is, -a).
//
// Note on overflow: for a == math.MinInt the mathematical negation is not
// representable as an int, and Go's two's-complement negation wraps back to
// math.MinInt. Callers relying on the involution property (Value(Value(x)) == x)
// must therefore avoid math.MinInt.
func Value(a int) int {
	return -a
}
