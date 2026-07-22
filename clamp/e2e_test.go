package clamp_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/clamp"
)

// TestE2EPageSizeClamping simulates a realistic use of Clamp: an API layer
// sanitizing client-supplied pagination page sizes into the server's allowed
// bounds before issuing a query. It drives the exported function through the
// public package boundary, exactly as a caller would.
func TestE2EPageSizeClamping(t *testing.T) {
	const (
		minPageSize = 10
		maxPageSize = 100
	)
	requested := []int{-5, 0, 10, 42, 100, 250}
	want := []int{10, 10, 10, 42, 100, 100}

	for i, r := range requested {
		got := clamp.Clamp(r, minPageSize, maxPageSize)
		if got != want[i] {
			t.Errorf("requested page size %d clamped to %d, want %d", r, got, want[i])
		}
	}
}
