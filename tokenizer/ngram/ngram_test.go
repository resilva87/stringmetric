package ngram

import (
	"reflect"
	"testing"
)

var ngramStringTests = []struct {
	title       string
	value       string
	N           int
	expected    []string
	expectedErr error
}{
	{
		title:       "should error for empty value",
		value:       "",
		N:           1,
		expected:    nil,
		expectedErr: ErrInvalidParams,
	},
	{
		title:       "should error for zero N",
		value:       "lorem ipsum",
		N:           0,
		expected:    nil,
		expectedErr: ErrInvalidParams,
	},
	{
		title:       "should error for negative N",
		value:       "lorem ipsum",
		N:           -3,
		expected:    nil,
		expectedErr: ErrInvalidParams,
	},
	{
		title:       "should error len(value) < N",
		value:       "lorem",
		N:           10,
		expected:    nil,
		expectedErr: ErrInvalidParams,
	},
	{
		title:       "should tokenize",
		value:       "lorem ipsum",
		N:           3,
		expected:    []string{"lor", "ore", "rem", "em ", "m i", " ip", "ips", "psu", "sum"},
		expectedErr: nil,
	},
	{
		title:       "should tokenize accented words",
		value:       "résumé",
		N:           2,
		expected:    []string{"ré", "és", "su", "um", "mé"},
		expectedErr: nil,
	},
}

func TestTokenize(t *testing.T) {
	for _, test := range ngramStringTests {
		t.Log(test.title)
		tokenizer := NewTokenizer(test.N)
		result, err := tokenizer.Tokenize(test.value)
		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("result doesn't match, expected %v, got: %v", test.expected, result)
		}
		if test.expectedErr != err {
			t.Errorf("error doesn't match, expected \"%v\", got: \"%v\"", test.expectedErr, err)
		}
	}
}
