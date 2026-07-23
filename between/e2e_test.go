package between_test

import (
	"testing"

	"github.com/hanahmily/tdd-sandbox/between"
)

// TestE2EOperatingRange simulates a real use case for the between package: a
// monitoring component screens a stream of sensor readings and accepts only
// those that fall within an inclusive acceptable operating range [lo, hi],
// rejecting readings on or beyond either edge. This exercises Between exactly
// as a downstream caller would, through the package's public import path.
func TestE2EOperatingRange(t *testing.T) {
	const (
		lo = 18
		hi = 24
	)

	readings := []int{17, 18, 21, 24, 25}
	want := []bool{false, true, true, true, false}

	var accepted []int
	for i, r := range readings {
		ok := between.Between(r, lo, hi)
		if ok != want[i] {
			t.Fatalf("reading[%d]=%d: Between(%d, %d, %d) = %v, want %v",
				i, r, r, lo, hi, ok, want[i])
		}
		if ok {
			accepted = append(accepted, r)
		}
	}

	wantAccepted := []int{18, 21, 24}
	if len(accepted) != len(wantAccepted) {
		t.Fatalf("accepted %v, want %v", accepted, wantAccepted)
	}
	for i, v := range wantAccepted {
		if accepted[i] != v {
			t.Fatalf("accepted %v, want %v", accepted, wantAccepted)
		}
	}
}
