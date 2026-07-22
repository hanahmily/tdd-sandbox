// Package e2e exercises the tally package the way a real downstream consumer
// would: importing the published module path and calling only its exported
// API. It is written RED-first and must fail until Total is implemented.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/tally"
)

// TestE2EOrderTotal simulates a realistic use case: a checkout service adds up
// the per-line subtotals (in cents) of a customer's shopping cart to produce
// the order total, including a discount line represented as a negative amount.
//
// The e2e gate runs `go test -run TestE2E ./...`, so this function must carry
// the TestE2E prefix.
func TestE2EOrderTotal(t *testing.T) {
	// Each entry is a cart line subtotal in cents. The negative line models a
	// coupon/discount applied at checkout.
	cart := map[string]int{
		"widget":   1299,  // $12.99
		"gadget":   4500,  // $45.00
		"shipping": 599,   // $5.99
		"coupon":   -1000, // -$10.00
	}

	lineItems := make([]int, 0, len(cart))
	for _, cents := range cart {
		lineItems = append(lineItems, cents)
	}

	const wantCents = 1299 + 4500 + 599 - 1000 // 5398 cents => $53.98

	gotCents := tally.Total(lineItems)
	if gotCents != wantCents {
		t.Fatalf("order total = %d cents, want %d cents ($%.2f)",
			gotCents, wantCents, float64(wantCents)/100)
	}
}

// TestE2EEmptyCart confirms the real-world edge case that an empty cart totals
// zero through the public API.
func TestE2EEmptyCart(t *testing.T) {
	if got := tally.Total(nil); got != 0 {
		t.Fatalf("empty cart total = %d, want 0", got)
	}
}

// TestE2ENetRefundTotal simulates a realistic account-adjustment run where the
// credits/refunds applied to an order exceed its charges, so the net movement
// is negative. It proves sign preservation survives all the way through the
// public API for a believable downstream use case, not only in unit tests.
func TestE2ENetRefundTotal(t *testing.T) {
	// Signed ledger entries in cents: one remaining charge against two refunds
	// (a returned item and a goodwill credit) that together exceed it.
	ledger := map[string]int{
		"restocking fee":  250,   // +$2.50 charge
		"item refund":     -4500, // -$45.00
		"goodwill credit": -1000, // -$10.00
	}

	entries := make([]int, 0, len(ledger))
	for _, cents := range ledger {
		entries = append(entries, cents)
	}

	const wantCents = 250 - 4500 - 1000 // -5250 cents => -$52.50

	gotCents := tally.Total(entries)
	if gotCents != wantCents {
		t.Fatalf("net ledger total = %d cents, want %d cents ($%.2f)",
			gotCents, wantCents, float64(wantCents)/100)
	}
}
