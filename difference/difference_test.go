package difference

import "testing"

// TestSub exercises Sub across positive, negative, and zero operands.
// It is RED against the not-implemented stub and turns green once Sub
// returns a - b.
func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive operands", a: 7, b: 3, want: 4},
		{name: "positive smaller minus larger", a: 3, b: 7, want: -4},
		{name: "negative operands", a: -5, b: -8, want: 3},
		{name: "negative minus positive", a: -5, b: 4, want: -9},
		{name: "positive minus negative", a: 6, b: -4, want: 10},
		{name: "zero minuend", a: 0, b: 5, want: -5},
		{name: "zero subtrahend", a: 5, b: 0, want: 5},
		{name: "both zero", a: 0, b: 0, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Sub(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
