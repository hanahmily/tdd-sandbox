package e2e_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2ESubLedgerBalance simulates a real use case: computing a running
// account balance by repeatedly subtracting debits from a starting balance.
// It consumes difference.Sub as an external caller would (black box) and
// covers positive, negative, and zero operands along the way.
//
// This is RED against the not-implemented stub and turns green once Sub
// returns a - b.
func TestE2ESubLedgerBalance(t *testing.T) {
	// Start with 100, apply a sequence of debits (some negative = credits,
	// one zero = no-op) and assert the balance after each step.
	balance := 100

	steps := []struct {
		debit       int
		wantBalance int
	}{
		{debit: 30, wantBalance: 70},   // positive operands
		{debit: 0, wantBalance: 70},    // zero subtrahend (no-op)
		{debit: -50, wantBalance: 120}, // negative debit acts as a credit
		{debit: 120, wantBalance: 0},   // drain to exactly zero
		{debit: 5, wantBalance: -5},    // overdraw into negative balance
	}

	for i, s := range steps {
		balance = difference.Sub(balance, s.debit)
		if balance != s.wantBalance {
			t.Fatalf("step %d: after debit %d, balance = %d, want %d",
				i, s.debit, balance, s.wantBalance)
		}
	}
}
