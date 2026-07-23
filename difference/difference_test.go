package difference

import "testing"

// TestSubPositive proves R1: Sub returns a - b for positive operands.
func TestSubPositive(t *testing.T) {
	if got := Sub(7, 3); got != 4 {
		t.Fatalf("Sub(7, 3) = %d, want 4", got)
	}
	// Positive result and positive (negative) result both exercised.
	if got := Sub(3, 7); got != -4 {
		t.Fatalf("Sub(3, 7) = %d, want -4", got)
	}
}

// TestSubNegative proves R2: Sub returns a - b when operands are negative.
func TestSubNegative(t *testing.T) {
	if got := Sub(-5, -2); got != -3 {
		t.Fatalf("Sub(-5, -2) = %d, want -3", got)
	}
	if got := Sub(-2, -5); got != 3 {
		t.Fatalf("Sub(-2, -5) = %d, want 3", got)
	}
}

// TestSubZero proves R3: Sub handles zero operands.
func TestSubZero(t *testing.T) {
	if got := Sub(0, 0); got != 0 {
		t.Fatalf("Sub(0, 0) = %d, want 0", got)
	}
	if got := Sub(5, 0); got != 5 {
		t.Fatalf("Sub(5, 0) = %d, want 5", got)
	}
	if got := Sub(0, 5); got != -5 {
		t.Fatalf("Sub(0, 5) = %d, want -5", got)
	}
}
