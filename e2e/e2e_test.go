package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EDifference proves R5: a real consumer can use difference.Sub across
// a package boundary to compute deltas over a sequence of meter readings —
// the net change between consecutive readings and the overall change.
func TestE2EDifference(t *testing.T) {
	readings := []int{10, 4, 4, -2}
	wantDeltas := []int{-6, 0, -6} // Sub(readings[i], readings[i-1])

	for i := 1; i < len(readings); i++ {
		got := difference.Sub(readings[i], readings[i-1])
		if got != wantDeltas[i-1] {
			t.Fatalf("delta %d: Sub(%d, %d) = %d, want %d",
				i-1, readings[i], readings[i-1], got, wantDeltas[i-1])
		}
	}

	if got := difference.Sub(readings[len(readings)-1], readings[0]); got != -12 {
		t.Fatalf("overall change: Sub(%d, %d) = %d, want -12",
			readings[len(readings)-1], readings[0], got)
	}
}
