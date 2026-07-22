package tally

// sum is the implementation seam behind the frozen Total interface.
//
// At contract time it ships as an unimplemented stub so the RED unit and
// end-to-end tests compile and fail (rather than breaking the build). This
// file is intentionally NOT in the protected list: the coder fills in the real
// summation here to turn the behavioural suite green.
//
// Implementation guidance for the coder:
//   - Accumulate the elements with native Go int arithmetic; the documented
//     contract adopts two's-complement wraparound on overflow (see tally.go and
//     TestTotalOverflowWrapsModulo), so do NOT add overflow detection, widening,
//     or an error path.
//   - Add no new exported identifiers — TestPublicAPISurfaceIsExactlyTotal locks
//     the surface to Total alone.
//   - Replace this comment with a present-tense description of the finished sum
//     when you implement it; do not leave stub/lifecycle wording in shipped code.
func sum(values []int) int {
	panic("tally: Total not implemented")
}
