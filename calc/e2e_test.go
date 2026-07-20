package calc_test

// End-to-end test: exercises the calc package the way a real external consumer
// would — importing it by its module path and composing Add into a
// higher-level workflow. This models a small "running ledger" that folds a
// stream of signed transactions (credits and debits) into a final balance
// using only the public API. It is RED against the contract stub and must pass
// once calc.Add is implemented. Its name matches the `TestE2E` gate pattern.

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

// runningBalance simulates a realistic client: it accumulates a sequence of
// signed amounts by repeatedly calling the public calc.Add, starting from a
// given opening balance. Debits are negative, credits are positive.
func runningBalance(opening int, transactions []int) int {
	balance := opening
	for _, amount := range transactions {
		balance = calc.Add(balance, amount)
	}
	return balance
}

func TestE2ELedger(t *testing.T) {
	cases := []struct {
		name         string
		opening      int
		transactions []int
		want         int
	}{
		{
			name:         "credits and debits settle positive",
			opening:      100,
			transactions: []int{50, -30, -20, 10}, // 100 +50 -30 -20 +10
			want:         110,
		},
		{
			name:         "overdraft goes negative",
			opening:      0,
			transactions: []int{-15, -35}, // 0 -15 -35
			want:         -50,
		},
		{
			name:         "empty statement keeps opening balance",
			opening:      42,
			transactions: nil,
			want:         42,
		},
		{
			name:         "transactions cancel to zero",
			opening:      0,
			transactions: []int{25, -25, 100, -100},
			want:         0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := runningBalance(tc.opening, tc.transactions); got != tc.want {
				t.Errorf("runningBalance(%d, %v) = %d, want %d",
					tc.opening, tc.transactions, got, tc.want)
			}
		})
	}
}
