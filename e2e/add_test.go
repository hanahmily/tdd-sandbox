package e2e_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/calc"
)

func TestAddEndToEndAccountBalance(t *testing.T) {
	openingBalance := 1_250
	deposit := 300
	withdrawal := -175

	balanceAfterDeposit := calc.Add(openingBalance, deposit)
	finalBalance := calc.Add(balanceAfterDeposit, withdrawal)

	const want = 1_375
	if finalBalance != want {
		t.Fatalf("final account balance = %d, want %d", finalBalance, want)
	}
}
