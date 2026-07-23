// Package e2e exercises the product package as an external consumer would,
// simulating a real use case rather than testing internals.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/product"
)

// TestE2EOrderLineTotal simulates an order-total calculation: a real caller
// multiplies a unit price by a quantity to produce a line total. It covers
// positive (a normal sale), negative (a refund/credit line), and zero (an
// out-of-stock or free line) operands end to end through the public API.
func TestE2EOrderLineTotal(t *testing.T) {
	lines := []struct {
		name      string
		unitPrice int
		quantity  int
		want      int
	}{
		{name: "standard sale", unitPrice: 250, quantity: 3, want: 750},
		{name: "refund credit line", unitPrice: -120, quantity: 2, want: -240},
		{name: "out of stock line", unitPrice: 999, quantity: 0, want: 0},
	}

	total := 0
	for _, ln := range lines {
		t.Run(ln.name, func(t *testing.T) {
			got := product.Mul(ln.unitPrice, ln.quantity)
			if got != ln.want {
				t.Errorf("product.Mul(%d, %d) = %d, want %d", ln.unitPrice, ln.quantity, got, ln.want)
			}
		})
		total += product.Mul(ln.unitPrice, ln.quantity)
	}

	const wantTotal = 750 - 240 + 0
	if total != wantTotal {
		t.Errorf("order total = %d, want %d", total, wantTotal)
	}
}
