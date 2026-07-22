package signint

import (
	"math"
	"testing"
)

// TestSign proves the three-way sign contract across the negative, zero, and
// positive cases, including the int boundary values.
func TestSign(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want int
	}{
		{"negative", -42, -1},
		{"minus_one", -1, -1},
		{"min_int", math.MinInt, -1},
		{"zero", 0, 0},
		{"positive", 42, 1},
		{"one", 1, 1},
		{"max_int", math.MaxInt, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sign(tc.in); got != tc.want {
				t.Fatalf("Sign(%d) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

// TestSignRangeIsMinusOneZeroOne proves Sign only ever returns -1, 0, or 1
// (never the raw input), across a sweep of representative values.
func TestSignRangeIsMinusOneZeroOne(t *testing.T) {
	inputs := []int{math.MinInt, -1000, -1, 0, 1, 1000, math.MaxInt}
	for _, in := range inputs {
		got := Sign(in)
		if got != -1 && got != 0 && got != 1 {
			t.Fatalf("Sign(%d) = %d, want one of {-1, 0, 1}", in, got)
		}
	}
}
