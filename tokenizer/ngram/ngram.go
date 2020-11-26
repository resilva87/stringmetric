package ngram

import (
	"errors"
)

type tokenizer interface {
	Tokenize(value string) ([]string, error)
	Tokenize2(value []byte) ([][]byte, error)
}

// Tokenizer is holder for parametrized n-gram tokenizer
type Tokenizer struct {
	N int
}

// ErrInvalidParams is thrown when trying to tokenize string with invalid params:
// N <= 0 or len(string) < N
var ErrInvalidParams error = errors.New("Invalid params to tokenize")

// NewTokenizer returns a new instance of n-gram tokenizer
func NewTokenizer(n int) *Tokenizer {
	return &Tokenizer{
		N: n,
	}
}

// Tokenize will break the string in tokens given a positive integer
// For instance, the word "lorem" tokenized in 3 results in ['lor', 'ore', 'rem']
func (t *Tokenizer) Tokenize(value string) ([]string, error) {
	internalRunes := []rune(value)
	vLen := len(internalRunes)
	if t.N <= 0 || vLen < t.N {
		return nil, ErrInvalidParams
	}
	var output []string

	for i := 0; i <= vLen-t.N; i++ {
		runes := internalRunes[i : i+t.N]
		output = append(output, string(runes))
	}
	return output, nil
}
