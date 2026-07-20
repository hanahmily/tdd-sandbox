// Package e2e exercises the calc package the way a real, external consumer
// would: it imports the published API and uses calc.Add to compute a running
// account balance over a ledger of transactions.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

// TestE2ELedgerRunningBalance simulates a real use case. Starting from a zero
// balance, it applies a sequence of deposits (positive) and withdrawals
// (negative), accumulating the balance exclusively through calc.Add, and
// checks the running balance after every transaction as well as the final
// balance.
func TestE2ELedgerRunningBalance(t *testing.T) {
	transactions := []int{100, -30, 0, -20, 50}
	wantRunning := []int{100, 70, 70, 50, 100}

	balance := 0
	for i, tx := range transactions {
		balance = calc.Add(balance, tx)
		if balance != wantRunning[i] {
			t.Fatalf("after transaction %d (%+d): balance = %d, want %d",
				i, tx, balance, wantRunning[i])
		}
	}

	const wantFinal = 100
	if balance != wantFinal {
		t.Fatalf("final balance = %d, want %d", balance, wantFinal)
	}
}

// TestE2ELedgerNetsToZero simulates draining an account back to zero, ensuring
// that a mix of positive and negative amounts cancels out exactly.
func TestE2ELedgerNetsToZero(t *testing.T) {
	transactions := []int{250, -100, -50, -100}

	balance := 0
	for _, tx := range transactions {
		balance = calc.Add(balance, tx)
	}

	if balance != 0 {
		t.Fatalf("final balance = %d, want 0", balance)
	}
}
