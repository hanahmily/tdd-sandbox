package calc_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

// TestLedgerBalanceE2E simulates a real consumer of the calc package: an
// account ledger that computes a running balance by folding calc.Add over a
// stream of transactions. Deposits are positive, withdrawals are negative and
// no-op entries are zero, so a single realistic scenario exercises the whole
// contract (positive, negative and zero operands) the way a caller would.
func TestLedgerBalanceE2E(t *testing.T) {
	type ledger struct {
		name         string
		transactions []int
		want         int
	}

	ledgers := []ledger{
		{
			name:         "deposits, withdrawals and no-ops net positive",
			transactions: []int{100, -30, 0, 50, -120, 0, 25},
			want:         25,
		},
		{
			name:         "overdraft ends negative",
			transactions: []int{40, -100, 0, 10},
			want:         -50,
		},
		{
			name:         "empty ledger stays at zero",
			transactions: nil,
			want:         0,
		},
	}

	for _, l := range ledgers {
		t.Run(l.name, func(t *testing.T) {
			balance := 0
			for _, amount := range l.transactions {
				balance = calc.Add(balance, amount)
			}
			if balance != l.want {
				t.Fatalf("final balance for %q = %d, want %d", l.name, balance, l.want)
			}
		})
	}
}

// TestReconciliationE2E checks that summing the same transactions in reverse
// order yields the same balance, mirroring how a real system might reconcile a
// ledger replayed from opposite ends.
func TestReconciliationE2E(t *testing.T) {
	transactions := []int{100, -30, 0, 50, -120, 0, 25}

	forward := 0
	for _, amount := range transactions {
		forward = calc.Add(forward, amount)
	}

	backward := 0
	for i := len(transactions) - 1; i >= 0; i-- {
		backward = calc.Add(backward, transactions[i])
	}

	if forward != backward {
		t.Fatalf("reconciliation mismatch: forward = %d, backward = %d", forward, backward)
	}
}
