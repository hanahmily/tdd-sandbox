package subtract

import "testing"

// TestDifference proves R1: Difference(a, b) returns a - b across a
// representative table of positive, negative, zero, and mixed inputs.
func TestDifference(t *testing.T) {
	cases := []struct {
		name    string
		a, b    int
		want    int
	}{
		{"both positive", 5, 3, 2},
		{"result negative", 3, 5, -2},
		{"identical operands", 7, 7, 0},
		{"subtract zero", 9, 0, 9},
		{"subtract from zero", 0, 4, -4},
		{"both negative", -5, -3, -2},
		{"mixed signs", -5, 3, -8},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Difference(tc.a, tc.b); got != tc.want {
				t.Errorf("Difference(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
