package stringmetric

import (
	"bytes"
	"testing"
)

var sameElementsTests = []struct {
	title    string
	a        []byte
	b        []byte
	expected bool
}{
	{
		title:    "should return true if both byte arrays are nil",
		a:        nil,
		b:        nil,
		expected: true,
	},
	{
		title:    "should return false if a is nil and b is not",
		a:        nil,
		b:        []byte{111},
		expected: false,
	},
	{
		title:    "should return false if b is nil and a is not",
		a:        []byte{111},
		b:        nil,
		expected: false,
	},
	{
		title:    "should return true if both byte arrays are empty",
		a:        []byte{},
		b:        []byte{},
		expected: true,
	},
	{
		title:    "should return false if a is empty and b is not",
		a:        []byte{},
		b:        []byte{111},
		expected: false,
	},
	{
		title:    "should return false if b is empty and a is not",
		a:        []byte{111},
		b:        []byte{},
		expected: false,
	},
	{
		title:    "should return false if lengths are different",
		a:        []byte{111},
		b:        []byte{111, 112},
		expected: false,
	},
	{
		title:    "should return false if elements are not in order",
		a:        []byte{112, 111},
		b:        []byte{111, 112},
		expected: false,
	},
	{
		title:    "should return true if elements are exactly the same",
		a:        []byte{111, 112},
		b:        []byte{111, 112},
		expected: true,
	},
}

func TestSameElements(t *testing.T) {
	for _, test := range sameElementsTests {
		t.Log(test.title)
		result := sameElements(test.a, test.b)
		if result != test.expected {
			t.Errorf("result doesn't match, expected %v, got: %v", test.expected, result)
		}
	}
}

var takeRightTests = []struct {
	title    string
	v        []byte
	size     int
	expected []byte
}{
	{
		title:    "should return empty array if size is zero",
		v:        []byte{111},
		size:     0,
		expected: []byte{},
	},
	{
		title:    "should return the whole array if size is greater than length",
		v:        []byte{111},
		size:     2,
		expected: []byte{111},
	},
	{
		title:    "should return an array from the n-th position to the end",
		v:        []byte{111, 112, 113},
		size:     2,
		expected: []byte{112, 113},
	},
}

func TestTakeRight(t *testing.T) {
	for _, test := range takeRightTests {
		t.Log(test.title)
		result := takeRight(test.v, test.size)
		if !bytes.Equal(result, test.expected) {
			t.Errorf("result doesn't match, expected %v, got: %v", test.expected, result)
		}
	}
}
