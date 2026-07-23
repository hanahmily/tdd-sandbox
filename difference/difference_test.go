package difference

import "testing"

// TestSub is the white-box unit test for Sub. It exercises the operand
// categories required by the milestone: positive, negative, and zero.
func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive operands", a: 7, b: 3, want: 4},
		{name: "positive smaller minus larger", a: 3, b: 7, want: -4},
		{name: "negative operands", a: -7, b: -3, want: -4},
		{name: "negative minus positive", a: -5, b: 4, want: -9},
		{name: "positive minus negative", a: 5, b: -4, want: 9},
		{name: "zero minuend", a: 0, b: 6, want: -6},
		{name: "zero subtrahend", a: 6, b: 0, want: 6},
		{name: "both zero", a: 0, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.a, tt.b); got != tt.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
