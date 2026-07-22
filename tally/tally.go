// Package tally provides helpers for summing collections of integers.
package tally

// Total returns the arithmetic sum of every element in values.
//
// Contract:
//   - An empty or nil slice yields 0 (the additive identity).
//   - The result is the sum of all elements, preserving sign, so a slice
//     containing negative values may produce a negative total.
//
// This is the approved, frozen public interface for the milestone. The
// signature must not change; the implementation is supplied separately via
// the sum seam so this file can remain locked while behaviour is filled in.
func Total(values []int) int {
	return sum(values)
}
