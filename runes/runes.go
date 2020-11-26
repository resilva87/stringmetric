package runes

import "unicode/utf8"

// MakeSlice lorem ipsum...
func MakeSlice(str string) []rune {
	var output []rune
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		output = append(output, r)
		str = str[size:]
	}
	return output
}
