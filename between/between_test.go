package between

import "testing"

// TestBetween exercises the inclusive-range semantics of Between across the
// five canonical positions relative to the range [min, max]: below, at the
// lower boundary, strictly inside, at the upper boundary, and above.
func TestBetween(t *testing.T) {
	tests := []struct {
		name  string
		value int
		min   int
		max   int
		want  bool
	}{
		{name: "below", value: 0, min: 1, max: 10, want: false},
		{name: "lower boundary", value: 1, min: 1, max: 10, want: true},
		{name: "inside", value: 5, min: 1, max: 10, want: true},
		{name: "upper boundary", value: 10, min: 1, max: 10, want: true},
		{name: "above", value: 11, min: 1, max: 10, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.value, tt.min, tt.max); got != tt.want {
				t.Errorf("Between(%d, %d, %d) = %v, want %v",
					tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
