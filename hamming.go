package stringmetric

// HammingMetric calculates the hamming distance between
// two strings
func HammingMetric(a string, b string) float64 {
	aRunes := []rune(a)
	bRunes := []rune(b)

	if len(aRunes) == 0 || len(bRunes) == 0 || len(aRunes) != len(bRunes) {
		return float64(-1)
	}
	if sameRunes(aRunes, bRunes) {
		return float64(0)
	}

	count := 0
	for i, aValue := range aRunes {
		bValue := bRunes[i]
		if aValue != bValue {
			count++
		}
	}

	return float64(count)
}
