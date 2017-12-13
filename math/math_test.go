package math

import "testing"

type testpair struct {
	values []float64
	expected float64
}

var tests = []testpair{
	{[]float64{1,2}, 1.5},
	{[]float64{-1,1}, 0},
	{[]float64{1,1,1,1,1,1}, 1},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error(
				"For ", pair.values,
					  "expected ", pair.expected,
					  "got ", v,
			)
		}
	}
}