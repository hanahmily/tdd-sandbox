// Package e2e exercises the calc package the way a real downstream consumer
// would: importing it through its public module path and using Add to solve an
// end-to-end problem, rather than testing a function in isolation.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

// runningBalance is a small realistic client of calc.Add. It walks a ledger of
// transactions — deposits are positive, withdrawals are negative, and a
// no-op adjustment is zero — accumulating a balance purely through calc.Add.
// This is the "real use case": accumulation over mixed-sign integers.
func runningBalance(transactions []int) int {
	balance := 0
	for _, txn := range transactions {
		balance = calc.Add(balance, txn)
	}
	return balance
}

// TestE2ELedger simulates a full account lifecycle: an opening deposit,
// several withdrawals that push the balance through zero and into the negative,
// a zero-value adjustment, and a recovering deposit. The final balance can only
// be correct if Add handles positive, negative, and zero operands correctly.
func TestE2ELedger(t *testing.T) {
	ledger := []int{
		100, // opening deposit
		-30, // withdrawal
		-70, // withdrawal that brings balance to exactly zero
		0,   // no-op adjustment
		-25, // overdraft into the negative
		50,  // recovering deposit
	}

	got := runningBalance(ledger)
	const want = 25 // 100 - 30 - 70 + 0 - 25 + 50

	if got != want {
		t.Fatalf("runningBalance(%v) = %d, want %d", ledger, got, want)
	}
}

// TestE2EEmptyLedger checks the boundary case: no transactions must leave
// the balance at zero, confirming the accumulator's identity behaviour.
func TestE2EEmptyLedger(t *testing.T) {
	if got := runningBalance(nil); got != 0 {
		t.Fatalf("runningBalance(nil) = %d, want 0", got)
	}
}

// TestE2EAllNegativeLedger confirms an account that only ever loses money
// accumulates a correct negative total.
func TestE2EAllNegativeLedger(t *testing.T) {
	got := runningBalance([]int{-5, -10, -15})
	const want = -30

	if got != want {
		t.Fatalf("runningBalance of all-negative ledger = %d, want %d", got, want)
	}
}
