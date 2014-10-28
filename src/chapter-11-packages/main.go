package main

import "fmt"
import m "chapter-11-packages/math"

func main() {
    xs := []float64{1, 2, 3, 4}
    avg := m.Average(xs)
    fmt.Println(avg)
    sml := m.Min(xs)
    fmt.Println(sml)
    grt := m.Max(xs)
    fmt.Println(grt)
}

// Problems
/*
We use packages to stay DRY. No need to reinvent the wheel.
*/

/*
An identifier doesn't start with a capital letter is hidden to the main routine.
*/

/*
If two packages have the same name you can use an alias to import the other one.
"import [alias] 'clone/package'"
*/
