package multiply_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/multiply"
)

// TestE2EProductCartTotal simulates a real use case: computing the line-item
// total for a shopping cart entry (unit price times quantity) through the
// public multiply.Product API, exactly as an external consumer would.
//
// It is named with the TestE2E prefix so `go test -run TestE2E ./...` matches
// it as the milestone's end-to-end gate.
func TestE2EProductCartTotal(t *testing.T) {
	type lineItem struct {
		name      string
		unitPrice int
		quantity  int
		wantTotal int
	}

	cart := []lineItem{
		{name: "widget", unitPrice: 250, quantity: 3, wantTotal: 750},
		{name: "gadget", unitPrice: 999, quantity: 0, wantTotal: 0},
		{name: "refund adjustment", unitPrice: -100, quantity: 2, wantTotal: -200},
	}

	grandTotal := 0
	for _, item := range cart {
		got := multiply.Product(item.unitPrice, item.quantity)
		if got != item.wantTotal {
			t.Fatalf("line total for %q: Product(%d, %d) = %d, want %d",
				item.name, item.unitPrice, item.quantity, got, item.wantTotal)
		}
		grandTotal += got
	}

	const wantGrandTotal = 750 + 0 - 200
	if grandTotal != wantGrandTotal {
		t.Fatalf("cart grand total = %d, want %d", grandTotal, wantGrandTotal)
	}
}
