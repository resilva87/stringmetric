package stringmetric

import (
	"github.com/resilva87/stringmetric/tokenizer/ngram"
)

// SorensenDiceMetric calculates strings similarity from two byte arrays
// using n-gram size
func SorensenDiceMetric(n int, a string, b string) float64 {
	if n <= 0 || len(a) == 0 || len(b) == 0 || len(a) < n || len(b) < n {
		return float64(0)
	}
	if sameElements([]byte(a), []byte(b)) {
		return float64(1)
	}
	tokenizer := ngram.NewTokenizer(n)
	aTokens, err := tokenizer.Tokenize(a)
	if err != nil {
		return float64(-1)
	}
	bTokens, err := tokenizer.Tokenize(b)
	if err != nil {
		return float64(-1)
	}

	intersections := hashIntersect(aTokens, bTokens)
	interLen := len(intersections)

	if interLen == 0 {
		return float64(0)
	}

	v := float64(2.0 * len(intersections))
	y := float64(len(aTokens) + len(bTokens))

	return v / y
}
