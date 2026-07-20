package calc_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

// TestAdd exercises the Add contract across positive, negative and zero
// operands. It is written against the public API only (external test package)
// so it verifies the approved interface, not any internal helper.
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "two positives", a: 2, b: 3, want: 5},
		{name: "positive and zero", a: 7, b: 0, want: 7},
		{name: "zero and positive", a: 0, b: 42, want: 42},
		{name: "zero and zero", a: 0, b: 0, want: 0},
		{name: "two negatives", a: -4, b: -6, want: -10},
		{name: "negative and zero", a: -9, b: 0, want: -9},
		{name: "negative and positive cancel", a: -5, b: 5, want: 0},
		{name: "negative plus larger positive", a: -3, b: 8, want: 5},
		{name: "positive plus larger negative", a: 10, b: -25, want: -15},
		{name: "large operands", a: 1_000_000, b: 2_345_678, want: 3_345_678},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calc.Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestAddZeroIdentity asserts that zero is the additive identity from both
// sides for a range of values.
func TestAddZeroIdentity(t *testing.T) {
	for _, x := range []int{0, 1, -1, 7, -128, 999} {
		if got := calc.Add(x, 0); got != x {
			t.Errorf("Add(%d, 0) = %d, want %d", x, got, x)
		}
		if got := calc.Add(0, x); got != x {
			t.Errorf("Add(0, %d) = %d, want %d", x, got, x)
		}
	}
}

// TestAddCommutative asserts Add(a, b) == Add(b, a).
func TestAddCommutative(t *testing.T) {
	pairs := [][2]int{{2, 3}, {-1, 9}, {0, -7}, {100, -250}, {-8, -8}}
	for _, p := range pairs {
		if calc.Add(p[0], p[1]) != calc.Add(p[1], p[0]) {
			t.Errorf("Add not commutative for %d, %d: %d != %d",
				p[0], p[1], calc.Add(p[0], p[1]), calc.Add(p[1], p[0]))
		}
	}
}
