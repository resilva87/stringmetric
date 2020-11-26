package stringmetric

import "reflect"

// sameElements checks if two byte arrays have the same elements in order
func sameElements(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		bValue := b[i]
		if v != bValue {
			return false
		}
	}
	return true
}

// sameRunes checks if two rune slices have the same elements in order
func sameRunes(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
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

// Taken from https://github.com/juliangruber/go-intersect
// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
func hashIntersect(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]bool)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = true
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}
