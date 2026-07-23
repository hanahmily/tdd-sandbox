package difference

import "testing"

// TestSub exercises Sub across positive, negative, and zero operands.
func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive_operands", a: 5, b: 3, want: 2},
		{name: "positive_result_from_larger_b", a: 3, b: 5, want: -2},
		{name: "negative_operands", a: -5, b: -3, want: -2},
		{name: "mixed_sign_operands", a: -5, b: 3, want: -8},
		{name: "zero_both", a: 0, b: 0, want: 0},
		{name: "zero_subtrahend", a: 4, b: 0, want: 4},
		{name: "zero_minuend", a: 0, b: 4, want: -4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sub(tc.a, tc.b); got != tc.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
