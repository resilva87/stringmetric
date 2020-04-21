package stringmetric

// sameElements checks if two byte arrays have the same elements in order
func sameElements(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range(a) {
		bValue := b[i]
		if v != bValue {
			return false
		}
	}
	return true
}

// takeRight takes a given number of elements from a byte array, starting
// from len(array) - size until len(array)
func takeRight(v []byte, size int) []byte {
	if size == 0 {
		return []byte{}
	}
	p := len(v) - size
	if p < 0 {
		p = 0
	}
	return v[p:]
}