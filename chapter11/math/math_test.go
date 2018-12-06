package math

import "testing"

type testpair struct {
	values  []float64
	Avarage float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAvarage(t *testing.T) {
	for _, pair := range tests {
		v := Avarage(pair.values)
		if v != pair.Avarage {
			t.Error(
				"For", pair.values,
				"expected", pair.Avarage,
				"got", v,
			)
		}
	}
}
