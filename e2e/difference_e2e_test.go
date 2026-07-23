// Package e2e exercises the difference package end to end, as an external
// consumer would, through its exported API only. These tests are RED until
// difference.Sub is implemented.
package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2ELedgerBalance simulates a real use case: an account ledger that
// starts from an opening balance and applies a sequence of debits, using
// difference.Sub to compute the running balance after each entry. It walks
// the balance down through a positive result, to exactly zero, and into a
// negative (overdrawn) result — covering positive, zero, and negative
// outcomes through the public interface.
func TestE2ELedgerBalance(t *testing.T) {
	const opening = 100

	debits := []int{40, 60, 25}
	// Running balance expected after each debit:
	//   100 - 40 = 60   (positive)
	//    60 - 60 =  0   (zero)
	//     0 - 25 = -25  (negative / overdrawn)
	wantAfter := []int{60, 0, -25}

	balance := opening
	for i, d := range debits {
		balance = difference.Sub(balance, d)
		if balance != wantAfter[i] {
			t.Fatalf("after debit %d of %d: balance = %d, want %d",
				i+1, d, balance, wantAfter[i])
		}
	}

	if balance != -25 {
		t.Errorf("final balance = %d, want -25", balance)
	}
}

// TestE2ESettleDifference simulates settling two parties' positions by
// computing the net difference between what each is owed, verifying the
// public Sub contract yields a symmetric, sign-correct result.
func TestE2ESettleDifference(t *testing.T) {
	alice, bob := 30, 45

	net := difference.Sub(alice, bob)
	if net != -15 {
		t.Fatalf("Sub(%d, %d) = %d, want -15", alice, bob, net)
	}

	// Reversing the operands must flip the sign exactly.
	if reversed := difference.Sub(bob, alice); reversed != 15 {
		t.Errorf("Sub(%d, %d) = %d, want 15", bob, alice, reversed)
	}

	// A party settling against itself nets to zero.
	if self := difference.Sub(alice, alice); self != 0 {
		t.Errorf("Sub(%d, %d) = %d, want 0", alice, alice, self)
	}
}
