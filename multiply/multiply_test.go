package multiply

import "testing"

// TestProduct is the RED unit test for Product. It exercises the core
// multiplication contract with a table of representative cases.
func TestProduct(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "positive", a: 3, b: 4, want: 12},
		{name: "with zero", a: 7, b: 0, want: 0},
		{name: "identity one", a: 1, b: 9, want: 9},
		{name: "negative times positive", a: -3, b: 5, want: -15},
		{name: "negative times negative", a: -6, b: -7, want: 42},
		{name: "large", a: 100000, b: 3, want: 300000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Product(tc.a, tc.b); got != tc.want {
				t.Fatalf("Product(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// TestProductCommutative verifies Product(a, b) == Product(b, a).
func TestProductCommutative(t *testing.T) {
	pairs := [][2]int{{2, 5}, {-4, 6}, {0, 8}, {-3, -9}}
	for _, p := range pairs {
		if Product(p[0], p[1]) != Product(p[1], p[0]) {
			t.Fatalf("Product not commutative for (%d, %d): %d != %d",
				p[0], p[1], Product(p[0], p[1]), Product(p[1], p[0]))
		}
	}
}
