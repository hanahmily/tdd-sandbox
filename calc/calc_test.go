package calc

import "testing"

// TestAdd exercises calc.Add across the required value classes: positive,
// negative, zero, and mixed-sign operands. These assertions are RED against
// the contract stub and must pass once Add returns a + b.
func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "two positives", a: 2, b: 3, want: 5},
		{name: "two negatives", a: -4, b: -6, want: -10},
		{name: "mixed signs", a: -7, b: 3, want: -4},
		{name: "zero and positive", a: 0, b: 9, want: 9},
		{name: "positive and zero", a: 9, b: 0, want: 9},
		{name: "both zero", a: 0, b: 0, want: 0},
		{name: "cancel to zero", a: 5, b: -5, want: 0},
		{name: "large operands", a: 1_000_000, b: 2_345_678, want: 3_345_678},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Add(tc.a, tc.b); got != tc.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// TestAddIsCommutative pins the commutativity property of addition, which the
// implementation must preserve for every pair of operands.
func TestAddIsCommutative(t *testing.T) {
	pairs := [][2]int{{1, 2}, {-3, 8}, {-5, -5}, {0, 42}}
	for _, p := range pairs {
		if forward, reverse := Add(p[0], p[1]), Add(p[1], p[0]); forward != reverse {
			t.Errorf("Add(%d, %d)=%d not commutative with Add(%d, %d)=%d",
				p[0], p[1], forward, p[1], p[0], reverse)
		}
	}
}

// TestAddZeroIdentity pins 0 as the additive identity element.
func TestAddZeroIdentity(t *testing.T) {
	for _, n := range []int{0, 1, -1, 12345, -98765} {
		if got := Add(n, 0); got != n {
			t.Errorf("Add(%d, 0) = %d, want %d", n, got, n)
		}
	}
}
