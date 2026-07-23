package difference_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestSub exercises Sub across positive, negative, and zero operands.
func TestSub(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive operands, positive result", 5, 3, 2},
		{"positive operands, negative result", 3, 5, -2},
		{"negative operands", -5, -3, -2},
		{"negative minus positive", -5, 3, -8},
		{"positive minus negative", 5, -3, 8},
		{"zero operands", 0, 0, 0},
		{"zero minuend", 0, 5, -5},
		{"zero subtrahend", 5, 0, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := difference.Sub(tt.a, tt.b); got != tt.want {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
