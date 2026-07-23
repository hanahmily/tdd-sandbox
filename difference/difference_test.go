package difference

import "testing"

// TestSub is the unit test for Sub. It covers positive, negative, and zero
// operands, and the non-commutative ordering of the difference.
func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive operands", a: 5, b: 3, want: 2},
		{name: "positive result smaller minuend", a: 3, b: 5, want: -2},
		{name: "negative operands", a: -5, b: -3, want: -2},
		{name: "mixed signs", a: -5, b: 3, want: -8},
		{name: "zero minuend", a: 0, b: 5, want: -5},
		{name: "zero subtrahend", a: 5, b: 0, want: 5},
		{name: "both zero", a: 0, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.a, tt.b); got != tt.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestSubAntiCommutative asserts Sub(a, b) == -Sub(b, a) for all sampled pairs.
func TestSubAntiCommutative(t *testing.T) {
	pairs := [][2]int{{5, 3}, {-5, -3}, {0, 7}, {-4, 9}}
	for _, p := range pairs {
		if got, mirror := Sub(p[0], p[1]), Sub(p[1], p[0]); got != -mirror {
			t.Errorf("Sub(%d, %d) = %d, but -Sub(%d, %d) = %d", p[0], p[1], got, p[1], p[0], -mirror)
		}
	}
}
