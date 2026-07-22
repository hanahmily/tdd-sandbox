package tally

// sum is the implementation seam behind the frozen Total interface.
//
// It ships as an unimplemented stub so that the RED unit and end-to-end tests
// compile and fail (rather than breaking the build). The coder implements the
// real summation here to turn the suite green; this file is intentionally NOT
// in the protected list.
func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}
