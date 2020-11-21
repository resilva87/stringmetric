package stringmetric

// longestCommonSequence calculates the longest common sequence of two byte arrays
// and returns a 3-column array which defines the length, row and column
func longestCommonSequence(a []byte, b []byte) [3]int {
	rows := len(a) + 1
	columns := len(b) + 1
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, columns)
	}
	lrc := [3]int{}
	for i, vA := range a {
		for j, vB := range b {
			if vA == vB {
				l := m[i][j] + 1
				m[i+1][j+1] = l
				if l > lrc[0] {
					lrc[0] = l
					lrc[1] = i + 1
					lrc[2] = j + 1
				}
			}
		}
	}
	return lrc
}

// commonSequences calculates all common sequences from two byte arrays.
// For instance, from the string "test1okay" and "test2okax" it should return
// a 2-row matrix, containing "test" and "oka" as common sequences.
func commonSequences(a []byte, b []byte) [][]byte {
	lcs := longestCommonSequence(a, b)
	if lcs[0] == 0 {
		return [][]byte{}
	}

	a1Size := lcs[1] - lcs[0]
	a1 := a[:a1Size]

	a2Size := len(a) - lcs[1]
	a2 := takeRight(a, a2Size)

	b1Size := lcs[2] - lcs[0]
	b1 := b[:b1Size]

	b2Size := len(b) - lcs[2]
	b2 := takeRight(b, b2Size)

	var v [][]byte
	r := a[a1Size:lcs[1]]
	v = append(v, r)
	v = append(v, commonSequences(a1, b1)...)
	v = append(v, commonSequences(a2, b2)...)
	return v
}

// RatcliffObershelpMetric calculates strings similarity from two strings
func RatcliffObershelpMetric(a string, b string) float64 {
	return RatcliffObershelpMetric2([]byte(a), []byte(b))
}

// RatcliffObershelpMetric2 calculates strings similarity from two byte arrays (chars)
func RatcliffObershelpMetric2(a []byte, b []byte) float64 {
	if a == nil || b == nil || len(a) == 0 || len(b) == 0 {
		return float64(0)
	}
	if sameElements(a, b) {
		return float64(1)
	}
	totalSize := 0
	for _, v := range commonSequences(a, b) {
		totalSize += len(v)
	}
	return float64(2.0*totalSize) / float64(len(a)+len(b))
}
