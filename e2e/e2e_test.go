// Package e2e exercises the difference package exactly as an external consumer
// would: importing it via its module path and calling the exported API.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EDifferenceSub simulates a real use case: a consumer computing running
// differences over a sequence of operations that span positive, negative, and
// zero operands. It is a RED end-to-end test — it fails until difference.Sub is
// implemented — and is discoverable by the e2e gate's `go test -run TestE2E`.
func TestE2EDifferenceSub(t *testing.T) {
	// A small ledger-style scenario: start from a balance and apply debits by
	// subtracting each amount, verifying the running balance at every step.
	steps := []struct {
		name    string
		balance int
		debit   int
		want    int
	}{
		{name: "debit from positive balance", balance: 100, debit: 40, want: 60},
		{name: "debit into negative balance", balance: 20, debit: 50, want: -30},
		{name: "debit of zero is a no-op", balance: -30, debit: 0, want: -30},
		{name: "debit a negative amount is a credit", balance: -30, debit: -30, want: 0},
	}

	balance := steps[0].balance
	for _, step := range steps {
		balance = difference.Sub(step.balance, step.debit)
		if balance != step.want {
			t.Fatalf("%s: Sub(%d, %d) = %d, want %d",
				step.name, step.balance, step.debit, balance, step.want)
		}
	}

	// Final independent assertion mirroring the documented contract.
	if got := difference.Sub(0, 0); got != 0 {
		t.Fatalf("Sub(0, 0) = %d, want 0", got)
	}
}
