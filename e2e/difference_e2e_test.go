package e2e_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EDifferenceSub simulates a real downstream consumer of the module's
// public surface: it imports the difference package and computes a running
// balance by repeatedly subtracting, covering positive, negative, and zero
// operands end to end.
func TestE2EDifferenceSub(t *testing.T) {
	// A small ledger: start at 10, apply a sequence of subtractions.
	balance := 10
	steps := []struct {
		amount int
		want   int // expected balance after subtracting amount
	}{
		{amount: 4, want: 6},   // positive operands: 10 - 4
		{amount: 0, want: 6},   // zero subtrahend: 6 - 0
		{amount: -3, want: 9},  // negative subtrahend raises balance: 6 - (-3)
		{amount: 9, want: 0},   // back to zero: 9 - 9
		{amount: 5, want: -5},  // crosses below zero: 0 - 5
	}

	for i, step := range steps {
		balance = difference.Sub(balance, step.amount)
		if balance != step.want {
			t.Fatalf("step %d: Sub(..., %d) produced balance %d, want %d",
				i, step.amount, balance, step.want)
		}
	}

	if balance != -5 {
		t.Errorf("final balance = %d, want -5", balance)
	}
}
