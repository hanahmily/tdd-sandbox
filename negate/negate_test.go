package negate

import (
	"math"
	"testing"
)

// TestValue proves R1, R2, R3: Value returns the arithmetic negation for
// positive, negative, and zero inputs.
func TestValue(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want int
	}{
		{"positive", 5, -5},
		{"negative", -7, 7},
		{"zero", 0, 0},
		{"one", 1, -1},
		{"large positive", 1_000_000, -1_000_000},
		{"large negative", -1_000_000, 1_000_000},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Value(tc.in); got != tc.want {
				t.Errorf("Value(%d) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

// TestValueInvolution proves R4: applying Value twice returns the original
// value for any int except math.MinInt.
func TestValueInvolution(t *testing.T) {
	for _, x := range []int{0, 1, -1, 42, -42, math.MaxInt, math.MinInt + 1} {
		if got := Value(Value(x)); got != x {
			t.Errorf("Value(Value(%d)) = %d, want %d", x, got, x)
		}
	}
}

// TestValueMinIntOverflow proves R5: the documented two's-complement overflow
// behavior for math.MinInt (negation wraps back to math.MinInt).
func TestValueMinIntOverflow(t *testing.T) {
	if got := Value(math.MinInt); got != math.MinInt {
		t.Errorf("Value(math.MinInt) = %d, want %d", got, math.MinInt)
	}
}
