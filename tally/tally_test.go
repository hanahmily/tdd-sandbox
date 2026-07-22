package tally

import "testing"

// TestTotal exercises the Total contract across the full range of inputs the
// milestone must satisfy. It is written RED-first: it fails against the stub
// implementation and must pass once Total is implemented.
func TestTotal(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "nil slice", values: nil, want: 0},
		{name: "empty slice", values: []int{}, want: 0},
		{name: "single element", values: []int{42}, want: 42},
		{name: "multiple positives", values: []int{1, 2, 3, 4, 5}, want: 15},
		{name: "all negatives", values: []int{-1, -2, -3}, want: -6},
		{name: "mixed signs and zero", values: []int{-5, 0, 10, -2, 7}, want: 10},
		{name: "single negative", values: []int{-99}, want: -99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Total(tt.values); got != tt.want {
				t.Errorf("Total(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}

// TestTotalDoesNotMutateInput guards the contract that Total is a pure read of
// its argument and never reorders or mutates the caller's slice.
func TestTotalDoesNotMutateInput(t *testing.T) {
	values := []int{3, 1, 2}
	snapshot := []int{3, 1, 2}

	_ = Total(values)

	for i := range values {
		if values[i] != snapshot[i] {
			t.Fatalf("Total mutated input at index %d: got %v, want %v", i, values, snapshot)
		}
	}
}
