package rekt

// max returns the lagrer int of the provided options
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// min returns the smaller int of the provided options
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// abs returns the posative counterpart of the given signed int
func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
