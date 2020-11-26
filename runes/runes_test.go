package runes

import (
	"reflect"
	"testing"
)

var makeSliceTests = []struct {
	title    string
	s1       string
	expected []rune
}{
	{
		title:    "should return nil if empty string",
		s1:       "",
		expected: nil,
	},
	{
		title:    "should return slice if string length is greater than zero",
		s1:       "renato",
		expected: []rune{'\u0072', '\u0065', '\u006E', '\u0061', '\u0074', '\u006F'},
	},
	{
		title:    "should return slice with accented words",
		s1:       "résumé",
		expected: []rune{'\u0072', '\u00E9', '\u0073', '\u0075', '\u006D', '\u00E9'},
	},
}

func TestMakeSlice(t *testing.T) {
	for _, test := range makeSliceTests {
		t.Log(test.title)
		result := MakeSlice(test.s1)
		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("result doesn't match, expected %v, got: %v", test.expected, result)
		}
	}
}
