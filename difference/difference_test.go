package difference

import "testing"

// TestSub exercises the Sub(a, b int) int contract across positive, negative,
// and zero operands. It is a RED test: it fails against the stub and passes
// once Sub correctly returns a - b.
func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive operands", a: 5, b: 3, want: 2},
		{name: "positive result smaller minuend", a: 3, b: 5, want: -2},
		{name: "negative operands", a: -5, b: -3, want: -2},
		{name: "mixed sign operands", a: -4, b: 6, want: -10},
		{name: "zero operands", a: 0, b: 0, want: 0},
		{name: "zero minuend", a: 0, b: 7, want: -7},
		{name: "zero subtrahend", a: 7, b: 0, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.a, tt.b); got != tt.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
