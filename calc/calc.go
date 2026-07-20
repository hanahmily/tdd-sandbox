// Package calc provides basic integer arithmetic operations.
//
// This file declares the approved public interface for the calc milestone.
// The behavior is intentionally left UNIMPLEMENTED: the accompanying unit and
// end-to-end tests are RED against this stub. The `implement` step replaces
// ONLY the function body (with `return a + b`), leaving the exported signature
// and doc contract below untouched.
package calc

// Add returns the sum of its two integer arguments.
//
// It must handle the full range of int inputs, including negative, zero, and
// mixed-sign operands, following the ordinary rules of integer addition:
//
//	Add(a, b) == a + b
//
// The operation is commutative (Add(a, b) == Add(b, a)) and 0 is the additive
// identity element (Add(a, 0) == a). Overflow follows native Go int semantics.
func Add(a, b int) int {
	panic("calc.Add: not implemented — the implement step must replace this body with `return a + b`")
}
