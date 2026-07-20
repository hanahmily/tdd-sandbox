package calc

import "testing"

// TestAdd exercises calc.Add across the required cases: two positives, negative
// operands, zero operands, and mixed signs. Every requirement in the milestone
// is represented by at least one row.
func TestAdd(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "two positives", a: 2, b: 3, want: 5},
		{name: "positive and zero", a: 7, b: 0, want: 7},
		{name: "zero and positive", a: 0, b: 7, want: 7},
		{name: "both zero", a: 0, b: 0, want: 0},
		{name: "two negatives", a: -4, b: -6, want: -10},
		{name: "negative and zero", a: -5, b: 0, want: -5},
		{name: "positive and negative cancel", a: 8, b: -8, want: 0},
		{name: "negative larger than positive", a: 3, b: -10, want: -7},
		{name: "positive larger than negative", a: -3, b: 10, want: 7},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Add(tc.a, tc.b); got != tc.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// TestAddIsCommutative documents that operand order does not affect the result,
// which must hold for all combinations of signs.
func TestAddIsCommutative(t *testing.T) {
	pairs := [][2]int{{5, 9}, {-5, 9}, {5, -9}, {-5, -9}, {0, 4}, {0, -4}}
	for _, p := range pairs {
		if Add(p[0], p[1]) != Add(p[1], p[0]) {
			t.Errorf("Add(%d, %d) != Add(%d, %d)", p[0], p[1], p[1], p[0])
		}
	}
}

// TestAddIdentity documents that zero is the additive identity for both
// positive and negative operands.
func TestAddIdentity(t *testing.T) {
	for _, n := range []int{0, 1, -1, 42, -42} {
		if got := Add(n, 0); got != n {
			t.Errorf("Add(%d, 0) = %d, want %d", n, got, n)
		}
		if got := Add(0, n); got != n {
			t.Errorf("Add(0, %d) = %d, want %d", n, got, n)
		}
	}
}
