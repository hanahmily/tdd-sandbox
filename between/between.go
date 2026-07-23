// Package between provides an inclusive numeric range check.
package between

// Between reports whether value lies within the inclusive range [min, max].
//
// It returns true when min <= value <= max, so both bounds are treated as
// part of the range. Callers are expected to pass min <= max; when min > max
// the range is empty and Between returns false for every value.
func Between(value, min, max int) bool {
	return value >= min && value <= max
}
