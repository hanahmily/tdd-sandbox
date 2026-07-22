package clamp

import "testing"

// TestClamp exercises the three range positions the contract must handle —
// below-range, in-range, and above-range — plus both inclusive boundaries.
func TestClamp(t *testing.T) {
	const (
		low  = 10
		high = 20
	)
	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "below range clamps up to low", value: 5, want: 10},
		{name: "in range returned unchanged", value: 15, want: 15},
		{name: "above range clamps down to high", value: 25, want: 20},
		{name: "equal to low bound", value: 10, want: 10},
		{name: "equal to high bound", value: 20, want: 20},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Clamp(tc.value, low, high); got != tc.want {
				t.Errorf("Clamp(%d, %d, %d) = %d, want %d", tc.value, low, high, got, tc.want)
			}
		})
	}
}
