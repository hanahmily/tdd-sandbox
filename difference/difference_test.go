package difference

import "testing"

// TestSub exercises the approved Sub(a, b int) int contract across positive,
// negative, and zero operands. It fails RED against the not-implemented stub.
func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		// Positive operands.
		{name: "positive_larger_minus_smaller", a: 5, b: 3, want: 2},
		{name: "positive_smaller_minus_larger", a: 3, b: 5, want: -2},
		{name: "positive_equal_operands", a: 7, b: 7, want: 0},

		// Negative operands.
		{name: "negative_minus_negative", a: -5, b: -3, want: -2},
		{name: "positive_minus_negative", a: 5, b: -3, want: 8},
		{name: "negative_minus_positive", a: -5, b: 3, want: -8},

		// Zero operands.
		{name: "zero_minus_zero", a: 0, b: 0, want: 0},
		{name: "value_minus_zero", a: 5, b: 0, want: 5},
		{name: "zero_minus_value", a: 0, b: 5, want: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.a, tt.b); got != tt.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
