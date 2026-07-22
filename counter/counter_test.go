package counter

import "testing"

// TestSum exercises the Sum contract directly against the package internals.
func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "positive values", values: []int{1, 2, 3}, want: 6},
		{name: "single value", values: []int{42}, want: 42},
		{name: "empty slice", values: []int{}, want: 0},
		{name: "nil slice", values: nil, want: 0},
		{name: "all negative values", values: []int{-1, -2, -3}, want: -6},
		{name: "mixed sign values", values: []int{-5, 10, -2, 1}, want: 4},
		{name: "values including zero", values: []int{0, 5, 0, -5, 0}, want: 0},
		{name: "cancelling values", values: []int{100, -100}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.values); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}
