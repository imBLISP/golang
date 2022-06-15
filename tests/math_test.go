package tests

import (
	"practice/math"
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		x := math.Average(pair.values)
		if x != pair.average {
			t.Error(
				"For", pair.values,
				"Expected", pair.average,
				"Got", x,
			)
		}
	}

	nums := []string{"vineet", "this"}
	nums = append(nums, "tesing")
}
