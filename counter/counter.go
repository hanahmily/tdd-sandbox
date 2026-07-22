// Package counter provides helpers for aggregating slices of integers.
package counter

// Sum returns the total of all values in the slice.
//
// Contract:
//   - A nil or empty slice sums to 0.
//   - Negative values are included in the total (they reduce it).
//   - No overflow handling beyond Go's native int wraparound is promised.
//
// This file declares the approved public interface only. The production
// behaviour is supplied separately by an unexported implementation
// (sumImpl); until that exists the package does not build, keeping the
// unit and end-to-end tests RED.
func Sum(values []int) int {
	return sumImpl(values)
}
