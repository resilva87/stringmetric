package runes

import "unicode/utf8"

// MakeSlice will turn a string into a slice of runes
func MakeSlice(str string) []rune {
	var output []rune
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		output = append(output, r)
		str = str[size:]
	}
	return output
}
