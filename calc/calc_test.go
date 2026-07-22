package calc

import "testing"

// TestMul verifies the core contract of Mul: it returns the arithmetic product
// of its two integer arguments across the sign and identity cases. These tests
// are RED until mul (mul.go) is implemented.
func TestMul(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "two positives", a: 6, b: 7, want: 42},
		{name: "multiply by zero", a: 5, b: 0, want: 0},
		{name: "positive times one", a: 99, b: 1, want: 99},
		{name: "negative times positive", a: -3, b: 4, want: -12},
		{name: "two negatives", a: -3, b: -4, want: 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.a, tt.b); got != tt.want {
				t.Errorf("Mul(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestMulIsCommutative checks that argument order does not change the result,
// a property callers rely on when computing quantity × price or price × quantity.
func TestMulIsCommutative(t *testing.T) {
	if Mul(3, 250) != Mul(250, 3) {
		t.Errorf("Mul is not commutative: Mul(3, 250) = %d, Mul(250, 3) = %d",
			Mul(3, 250), Mul(250, 3))
	}
}
