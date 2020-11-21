package stringmetric

import (
	"testing"
)

// TODO all tests are bigrams, should test for other values as well
var sorensenDiceMetricTests = []struct {
	title    string
	N        int
	s1       string
	s2       string
	expected float64
}{
	{
		title:    "should return no similarity if s1 is empty",
		N:        1,
		s1:       "",
		s2:       "hello",
		expected: float64(0),
	},
	{
		title:    "should return no similarity if s2 is empty",
		N:        1,
		s1:       "hello",
		s2:       "",
		expected: float64(0),
	},
	{
		title:    "should return full similarity if s1 and s2 have the same elements",
		N:        2,
		s1:       "hello",
		s2:       "hello",
		expected: float64(1),
	},
	{
		title:    "should compute similarity with words with punctuation",
		N:        2,
		s1:       "night!",
		s2:       "natch?",
		expected: float64(0),
	},
	{
		title:    "should compute similarity for accented words",
		N:        2,
		s1:       "résumé",
		s2:       "resumé",
		expected: float64(0.6),
	},
}

func TestSorensenDiceMetric(t *testing.T) {
	for _, test := range sorensenDiceMetricTests {
		t.Log(test.title)
		result := SorensenDiceMetric(test.N, test.s1, test.s2)
		if result != test.expected {
			t.Errorf("result doesn't match, expected %v, got: %v", test.expected, result)
		}
	}
}
