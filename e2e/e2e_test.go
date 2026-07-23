package e2e

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/difference"
)

// TestE2EDifferenceBalance exercises Sub the way a real caller would: it walks a
// series of meter readings and asserts the net change between consecutive
// readings, exercising positive, negative, and zero deltas end to end through
// the exported package boundary.
func TestE2EDifferenceBalance(t *testing.T) {
	// Sequential readings; each step's expected delta is reading[i] - reading[i-1].
	readings := []int{100, 130, 130, 90, 90, 0, -25}
	wantDeltas := []int{30, 0, -40, 0, -90, -25} // len == len(readings)-1

	if len(wantDeltas) != len(readings)-1 {
		t.Fatalf("test setup error: %d deltas for %d readings", len(wantDeltas), len(readings))
	}

	for i := 1; i < len(readings); i++ {
		got := difference.Sub(readings[i], readings[i-1])
		want := wantDeltas[i-1]
		if got != want {
			t.Errorf("delta between reading[%d]=%d and reading[%d]=%d: Sub = %d, want %d",
				i, readings[i], i-1, readings[i-1], got, want)
		}
	}

	// The sum of all consecutive deltas must equal the overall difference
	// (telescoping), a property a real consumer relies on.
	var sum int
	for i := 1; i < len(readings); i++ {
		sum += difference.Sub(readings[i], readings[i-1])
	}
	if overall := difference.Sub(readings[len(readings)-1], readings[0]); sum != overall {
		t.Errorf("telescoping failed: sum of deltas = %d, overall Sub = %d", sum, overall)
	}
}
