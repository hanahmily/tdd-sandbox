package product

import "testing"

// TestMul exercises the Mul contract across positive, negative, and zero
// operands. It fails against the stub and passes once Mul multiplies.
func TestMul(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive operands", a: 6, b: 7, want: 42},
		{name: "positive by one", a: 1, b: 9, want: 9},
		{name: "negative by positive", a: -4, b: 5, want: -20},
		{name: "negative by negative", a: -3, b: -8, want: 24},
		{name: "zero left operand", a: 0, b: 12, want: 0},
		{name: "zero right operand", a: 15, b: 0, want: 0},
		{name: "both zero", a: 0, b: 0, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Mul(tc.a, tc.b); got != tc.want {
				t.Errorf("Mul(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// TestMulCommutative documents that multiplication is order-independent.
func TestMulCommutative(t *testing.T) {
	if Mul(-7, 3) != Mul(3, -7) {
		t.Errorf("Mul is not commutative: Mul(-7,3)=%d, Mul(3,-7)=%d", Mul(-7, 3), Mul(3, -7))
	}
}
