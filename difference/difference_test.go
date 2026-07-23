package difference

import "testing"

// TestSub proves the Sub contract across the required operand classes:
// positive, negative, and zero. It is RED against the placeholder stub.
func TestSub(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"positive operands", 7, 3, 4},
		{"positive minus larger is negative", 3, 8, -5},
		{"negative operands", -5, -2, -3},
		{"negative minus positive", -4, 6, -10},
		{"zero operands", 0, 0, 0},
		{"subtract zero is identity", 5, 0, 5},
		{"zero minus positive", 0, 9, -9},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sub(tc.a, tc.b); got != tc.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
