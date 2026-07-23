package difference

import "testing"

// TestSub exercises the full behavioral contract of Sub across positive,
// negative, and zero operands. It is a white-box (internal package) unit
// test and is RED until difference.Sub is implemented.
func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive operands", 9, 4, 5},
		{"positive result negative", 4, 9, -5},
		{"negative operands", -3, -8, 5},
		{"mixed signs", -3, 5, -8},
		{"zero subtrahend identity", 7, 0, 7},
		{"zero minuend", 0, 7, -7},
		{"both zero", 0, 0, 0},
		{"equal operands", 6, 6, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sub(tc.a, tc.b); got != tc.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
