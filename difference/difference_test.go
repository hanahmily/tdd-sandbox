package difference

import "testing"

// TestSubPositive proves R1/R2: Sub exists with the approved signature and
// returns a - b for positive operands.
func TestSubPositive(t *testing.T) {
	if got := Sub(7, 3); got != 4 {
		t.Fatalf("Sub(7, 3) = %d, want 4", got)
	}
	if got := Sub(3, 7); got != -4 {
		t.Fatalf("Sub(3, 7) = %d, want -4", got)
	}
}

// TestSubNegative proves R3: Sub handles negative operands and negative
// results.
func TestSubNegative(t *testing.T) {
	if got := Sub(-5, -2); got != -3 {
		t.Fatalf("Sub(-5, -2) = %d, want -3", got)
	}
	if got := Sub(-5, 2); got != -7 {
		t.Fatalf("Sub(-5, 2) = %d, want -7", got)
	}
}

// TestSubZero proves R4: Sub handles zero operands and zero results.
func TestSubZero(t *testing.T) {
	if got := Sub(0, 0); got != 0 {
		t.Fatalf("Sub(0, 0) = %d, want 0", got)
	}
	if got := Sub(5, 0); got != 5 {
		t.Fatalf("Sub(5, 0) = %d, want 5", got)
	}
	if got := Sub(5, 5); got != 0 {
		t.Fatalf("Sub(5, 5) = %d, want 0", got)
	}
}
