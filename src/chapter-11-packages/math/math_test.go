// Part of testing chapter
package math

import "testing"

// Test Average
func TestAverage(t *testing.T) {
    var v float64
    v = Average([]float64{1,2})
    if v != 1.5 {
        t.Error("Expected 1.5, got ", v)
    }
}

// Test Max
func TestMax(t *testing.T) {
    var v float64
    v = Max([]float64{17,34,1})
    if v != 34 {
        t.Error("Expected 34, got ", v)
    }
}

//Test Min
func TestMin(t *testing.T) {
    var v float64
    v = Min([]float64{17,34,1})
    if v != 1 {
        t.Error("Expected 1, got ", v)
    }
}
