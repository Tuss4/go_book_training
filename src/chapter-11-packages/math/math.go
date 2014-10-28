//Part of the Chapter 11 problems
package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
    total := float64(0)
    if len(xs) == 0 {
        return total
    }
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Finds the smallest number in a series of numbers
func Min(xs []float64) float64 {
    if len(xs) == 0 {
        return float64(0)
    }
    i := xs[0]
    for _, x := range xs {
        if i > x {
            i = x
        }
    }
    return i
}

// Finds the greatest number in a series of numbers
func Max(xs []float64) float64 {
    if len(xs) == 0 {
        return float64(0)
    }
    i := float64(0)
    for _, x := range xs {
        if i < x {
            i = x
        }
    }
    return i
}
