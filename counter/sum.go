package counter

// sumImpl returns the total of values using Go's native int arithmetic.
func sumImpl(values []int) int {
	var total int
	for _, value := range values {
		total += value
	}
	return total
}
