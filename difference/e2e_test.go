package difference_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2ELedgerReconciliation simulates a real use case for Sub: reconciling
// an account ledger. Starting from an opening balance, each debit is applied
// by subtracting it from the running balance, and the closing balance is
// verified. It then drives the balance negative to model an overdraft,
// exercising positive, zero, and negative results end to end.
func TestE2ELedgerReconciliation(t *testing.T) {
	const opening = 100
	debits := []int{20, 35, 45}

	balance := opening
	for _, debit := range debits {
		balance = difference.Sub(balance, debit)
	}

	if balance != 0 {
		t.Fatalf("closing balance = %d, want 0", balance)
	}

	// One more debit pushes the account into overdraft (negative balance).
	if got := difference.Sub(balance, 10); got != -10 {
		t.Fatalf("overdraft balance = %d, want -10", got)
	}
}
