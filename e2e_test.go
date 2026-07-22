package signint

import (
	"testing"
)

// TestE2ETrendDetection simulates a real use case for Sign: classifying the
// trend between consecutive readings in a time series as falling (-1), flat
// (0), or rising (1) by taking the sign of each delta. This exercises Sign
// exactly as a downstream caller would.
func TestE2ETrendDetection(t *testing.T) {
	readings := []int{10, 10, 12, 7, 7, 20}
	want := []int{0, 1, -1, 0, 1}

	got := make([]int, 0, len(readings)-1)
	for i := 1; i < len(readings); i++ {
		got = append(got, Sign(readings[i]-readings[i-1]))
	}

	if len(got) != len(want) {
		t.Fatalf("trend length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("trend[%d] = %d, want %d (readings %d -> %d)",
				i, got[i], want[i], readings[i], readings[i+1])
		}
	}
}
