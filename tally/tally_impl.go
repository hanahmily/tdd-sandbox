package tally

// sum accumulates values using native Go int arithmetic.
func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}

	return total
}
