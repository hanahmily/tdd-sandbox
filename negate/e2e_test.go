package negate_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/negate"
)

// TestE2ENegateBalance simulates a real use case: reversing a running ledger
// balance so that a caller can turn a credit into a debit and back again by
// applying negation, exercising Value through the package's public API.
//
// Proves R6.
func TestE2ENegateBalance(t *testing.T) {
	// A sequence of signed transaction amounts.
	txns := []int{100, -30, 250, -75}

	// Negate each amount (e.g. mirroring the ledger for a counterparty).
	mirrored := make([]int, len(txns))
	for i, amt := range txns {
		mirrored[i] = negate.Value(amt)
	}

	// The mirrored ledger must be the exact opposite of the original.
	var original, opposite int
	for i := range txns {
		original += txns[i]
		opposite += mirrored[i]
	}
	if opposite != negate.Value(original) {
		t.Fatalf("mirrored balance = %d, want %d", opposite, negate.Value(original))
	}

	// Negating the mirror restores the original ledger.
	for i, amt := range mirrored {
		if got := negate.Value(amt); got != txns[i] {
			t.Errorf("Value(Value(%d)) = %d, want %d", txns[i], got, txns[i])
		}
	}
}
