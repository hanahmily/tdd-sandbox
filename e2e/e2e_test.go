// Package e2e exercises the difference package end to end, the way a real
// consumer would: it imports the public API by its module path and drives a
// realistic scenario built entirely on difference.Sub.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EBudgetTracker simulates a real use case: a spending tracker that
// starts from an opening balance and repeatedly applies difference.Sub to
// deduct each expense, then verifies the running and final balances. Because
// the scenario chains positive, negative (a refund modeled as subtracting a
// negative amount), and zero deltas through the public API, it proves the
// contract works in composition, not just in isolation. It fails RED until
// difference.Sub is implemented.
func TestE2EBudgetTracker(t *testing.T) {
	const opening = 100

	// name is descriptive; delta is the amount subtracted from the balance.
	// A negative delta models a refund (subtracting a negative adds money).
	steps := []struct {
		name        string
		delta       int
		wantBalance int
	}{
		{name: "groceries", delta: 30, wantBalance: 70},
		{name: "no-op charge", delta: 0, wantBalance: 70},
		{name: "rent", delta: 90, wantBalance: -20}, // balance goes negative (overdraft)
		{name: "refund", delta: -50, wantBalance: 30},
	}

	balance := opening
	for _, step := range steps {
		balance = difference.Sub(balance, step.delta)
		if balance != step.wantBalance {
			t.Fatalf("after %q: balance = %d, want %d", step.name, balance, step.wantBalance)
		}
	}

	const wantFinal = 30
	if balance != wantFinal {
		t.Fatalf("final balance = %d, want %d", balance, wantFinal)
	}
}
