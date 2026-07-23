package difference_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EDifferenceSettlement is the black-box end-to-end test. It imports the
// package the way a real consumer would (through the module path) and drives a
// realistic scenario: settling a running balance by repeatedly subtracting
// amounts. It relies only on the exported Sub API and covers positive,
// negative, and zero operands along the way.
func TestE2EDifferenceSettlement(t *testing.T) {
	// A ledger settles by subtracting each charge from the opening balance.
	balance := 100

	steps := []struct {
		charge int
		want   int // expected running balance after subtracting charge
	}{
		{charge: 30, want: 70},   // positive operands
		{charge: 0, want: 70},    // zero subtrahend leaves balance unchanged
		{charge: 90, want: -20},  // overdraft: result goes negative
		{charge: -20, want: 0},   // negative subtrahend (a refund) restores to zero
	}

	for i, step := range steps {
		balance = difference.Sub(balance, step.charge)
		if balance != step.want {
			t.Fatalf("step %d: after Sub with charge %d, balance = %d, want %d",
				i, step.charge, balance, step.want)
		}
	}

	// Final settlement: subtracting a value from itself yields zero.
	if got := difference.Sub(balance, balance); got != 0 {
		t.Errorf("final settlement Sub(%d, %d) = %d, want 0", balance, balance, got)
	}
}
