// Part of testing chapter
package math

import "testing"

type testpair struct {
    values []float64
    average float64
}

var tests = []testpair{
    { []float64{1,2}, 1.5},
    { []float64{1,1,1,1,1,1}, 1},
    { []float64{-1,1}, 0},
    { []float64{}, 0},
}

type hiTestPair struct {
    values []float64
    max float64
}

var hiTests = []hiTestPair{
    { []float64{1,2, 54}, 54},
    { []float64{0.000000001, 0.00001}, 0.00001},
    { []float64{-1,1}, 1},
    { []float64{}, 0},
}

type loTestPair struct {
    values []float64
    min float64
}

var loTests = []loTestPair{
    { []float64{1,2, 54}, 1},
    { []float64{0.000000001, 0.00001}, 0.000000001},
    { []float64{-1,1}, -1},
    { []float64{}, 0},
}
// Test Average
func TestAverage(t *testing.T) {
    for _, pair := range tests {
        v := Average(pair.values)
        if v != pair.average {
            t.Error(
                "For", pair.values,
                "expected", pair.average,
                "got", v,
            )
        }
    }
}

// Test Max
func TestMax(t *testing.T) {
    for _, pair := range hiTests {
        v := Max(pair.values)
        if v != pair.max {
            t.Error(
                "For", pair.values,
                "expected", pair.max,
                "got", v,
            )
        }
    }
}

//Test Min
func TestMin(t *testing.T) {
    for _, pair := range loTests {
        v := Min(pair.values)
        if v != pair.min {
            t.Error(
                "For", pair.values,
                "expected", pair.min,
                "got", v,
            )
        }
    }
}
