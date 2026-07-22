// Package e2e exercises the calc package the way a real consumer would,
// importing it across a package boundary rather than testing internals.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

// TestE2EOrderTotal simulates a real use case: a checkout computing the total
// cost of a shopping cart, where each line's cost is quantity × unit price
// (in cents) via calc.Mul, then the lines are summed into an order total.
//
// This is the end-to-end contract for the milestone and is RED until Mul is
// implemented.
func TestE2EOrderTotal(t *testing.T) {
	type line struct {
		name      string
		quantity  int
		unitPrice int // in cents
	}

	cart := []line{
		{name: "widget", quantity: 3, unitPrice: 250}, // 750
		{name: "gadget", quantity: 2, unitPrice: 400}, // 800
	}

	total := 0
	for _, l := range cart {
		lineCost := calc.Mul(l.quantity, l.unitPrice)
		total += lineCost
	}

	const want = 1550 // 3*250 + 2*400
	if total != want {
		t.Fatalf("order total = %d cents, want %d cents", total, want)
	}
}
