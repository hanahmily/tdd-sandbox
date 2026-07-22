// Package clamp provides a generic helper for constraining ordered
// integer values to an inclusive range.
package clamp

// Integer is the set of integer types [Clamp] operates on. It admits every
// built-in signed and unsigned integer kind, plus any named type whose
// underlying type is one of them.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Clamp constrains value to the inclusive range [low, high] and returns the
// result: low when value < low, high when value > high, and value itself when
// it already lies within the range (including when it equals either bound).
//
// The bounds must satisfy low <= high; the result of calling Clamp with
// low > high is undefined.
func Clamp[T Integer](value, low, high T) T {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}
