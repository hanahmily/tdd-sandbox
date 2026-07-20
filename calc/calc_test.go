package calc

import "testing"

// TestAdd covers the core requirement: Add returns the sum of its two
// arguments, including negative and zero operands.
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"two positives", 2, 3, 5},
		{"positive and zero", 5, 0, 5},
		{"zero and positive", 0, 7, 7},
		{"both zero", 0, 0, 0},
		{"two negatives", -4, -6, -10},
		{"negative and positive, sum positive", -3, 10, 7},
		{"negative and positive, sum negative", -10, 3, -7},
		{"positive and negative, sum zero", 8, -8, 0},
		{"negative and zero", -5, 0, -5},
		{"large positives", 1_000_000, 2_000_000, 3_000_000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestAddIsCommutative asserts that Add(a, b) == Add(b, a) for a range of
// operand sign combinations.
func TestAddIsCommutative(t *testing.T) {
	cases := [][2]int{{2, 3}, {-5, 8}, {0, -4}, {-7, -1}, {6, 0}}
	for _, c := range cases {
		if got, rev := Add(c[0], c[1]), Add(c[1], c[0]); got != rev {
			t.Errorf("Add(%d, %d) = %d but Add(%d, %d) = %d: not commutative",
				c[0], c[1], got, c[1], c[0], rev)
		}
	}
}
