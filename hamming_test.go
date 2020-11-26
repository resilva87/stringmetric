package stringmetric

import "testing"

var hammingMetrics = []struct {
	title    string
	s1       string
	s2       string
	expected float64
}{
	{
		title:    "should return no similarity if s1 is empty",
		s1:       "",
		s2:       "hello",
		expected: float64(-1),
	},
	{
		title:    "should return no similarity if s2 is empty",
		s1:       "hello",
		s2:       "",
		expected: float64(-1),
	},
	{
		title:    "should return full similarity if s1 and s2 have the same elements",
		s1:       "hello",
		s2:       "hello",
		expected: float64(0),
	},
	{
		title:    "should compute similarity with words with punctuation",
		s1:       "hello!",
		s2:       "helka?",
		expected: float64(3),
	},
	{
		title:    "should compute similarity for accented words",
		s1:       "résumé",
		s2:       "resumé",
		expected: float64(1),
	},
	{
		title:    "should compute exact similarity for the same word",
		s1:       "résumé",
		s2:       "résumé",
		expected: float64(0),
	},
}

func TestHammingMetric(t *testing.T) {
	for _, test := range hammingMetrics {
		t.Log(test.title)
		result := HammingMetric(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("result doesn't match, expected %v, got: %v", test.expected, result)
		}
	}
}
