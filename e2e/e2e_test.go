// Package e2e exercises the difference package the way a real caller would:
// as an external consumer importing only the exported surface.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EDifferenceLedger proves R4: a realistic ledger scenario that reduces a
// starting balance by a series of entries using difference.Sub, exercising
// positive, negative, and zero operands end to end through the public API.
func TestE2EDifferenceLedger(t *testing.T) {
	// A simple expense ledger: start with a budget and subtract each entry.
	// Entries include a normal debit (positive), a refund (negative debit,
	// which increases the balance), and a no-op (zero).
	balance := 100
	entries := []int{40, 0, -10, 25}

	for _, e := range entries {
		balance = difference.Sub(balance, e)
	}

	// 100 - 40 - 0 - (-10) - 25 = 45
	if balance != 45 {
		t.Fatalf("final ledger balance = %d, want 45", balance)
	}

	// The final settlement: the difference between the starting budget and the
	// remaining balance is the net spend.
	netSpend := difference.Sub(100, balance)
	if netSpend != 55 {
		t.Fatalf("net spend = %d, want 55", netSpend)
	}
}
