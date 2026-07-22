package counter_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/counter"
)

// TestE2ECounterLedgerBalance simulates a real use case: an accounting
// consumer that tallies a day's worth of transactions per account into a net
// balance. Deposits are positive amounts, withdrawals are negative, and an
// account with no activity settles at zero.
//
// It consumes the counter package purely through its exported API, the way a
// downstream user would, exercising the empty, negative, and mixed-sign paths
// called out by the milestone.
func TestE2ECounterLedgerBalance(t *testing.T) {
	type account struct {
		name         string
		transactions []int
		wantBalance  int
	}

	ledger := []account{
		{name: "active account", transactions: []int{100, -30, -20, 50}, wantBalance: 100},
		{name: "overdrawn account", transactions: []int{50, -200}, wantBalance: -150},
		{name: "deposit only", transactions: []int{10, 20, 30}, wantBalance: 60},
		{name: "dormant account", transactions: []int{}, wantBalance: 0},
	}

	var grandTotal int
	for _, acct := range ledger {
		balance := counter.Sum(acct.transactions)
		if balance != acct.wantBalance {
			t.Errorf("account %q: counter.Sum(%v) = %d, want %d",
				acct.name, acct.transactions, balance, acct.wantBalance)
		}
		grandTotal += balance
	}

	// The whole ledger nets out to the sum of every per-account balance.
	if wantGrandTotal := 10; grandTotal != wantGrandTotal {
		t.Errorf("ledger grand total = %d, want %d", grandTotal, wantGrandTotal)
	}
}
