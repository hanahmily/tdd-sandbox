// Package tally provides helpers for summing collections of integers.
package tally

// Total returns the arithmetic sum of every element in values.
//
// Contract:
//   - An empty or nil slice yields 0 (the additive identity).
//   - The result is the sum of all elements, preserving sign, so a slice
//     containing negative values may produce a negative total.
//   - Accumulation uses native Go int arithmetic. If the true mathematical
//     sum is not representable in an int on the target platform, Total wraps
//     modulo 2^bits (two's-complement wraparound) rather than panicking or
//     reporting an error. int is architecture-sized, so callers that need an
//     exact result must ensure the sum fits the target's int range; a value
//     that fits on a 64-bit build may wrap on a 32-bit build. This wraparound
//     is the defined, tested behaviour, not undefined behaviour.
//
// This is the approved, frozen public interface for the milestone. The
// signature must not change; the implementation is supplied separately via
// the sum seam so this file can remain locked while behaviour is filled in.
func Total(values []int) int {
	return sum(values)
}
