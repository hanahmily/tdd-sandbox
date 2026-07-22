// Package e2e exercises the subtract package as an external consumer,
// simulating a real use case: computing a running account balance by
// applying a sequence of debits with subtract.Difference.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/subtract"
)

// TestE2ESubtractRunningBalance proves R2: an external caller can use
// subtract.Difference to compute a running balance by repeatedly
// subtracting debits from a starting balance, arriving at the correct
// final total.
func TestE2ESubtractRunningBalance(t *testing.T) {
	balance := 100
	debits := []int{10, 25, 5, 40}
	want := 20 // 100 - 10 - 25 - 5 - 40

	for _, d := range debits {
		balance = subtract.Difference(balance, d)
	}

	if balance != want {
		t.Errorf("running balance after debits %v = %d, want %d", debits, balance, want)
	}
}
