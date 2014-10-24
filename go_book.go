package main

import (
    "fmt"
)

func main() {
    fmt.Println(true && false)
    fmt.Println(true && true)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)

    var a, b byte
    a = 255
    fmt.Printf("%v, %T, %b\n", a, a, a)
    fmt.Println(a & 1)
    b = 254
    if (b & 1 == 0) {
        fmt.Println("Yurrp")
    } else {
        fmt.Println("Naaaahhh")
    }
    var c int
    c = 32132*42452
    fmt.Printf("32132 x 42452 = %v\n", c)
    fmt.Println(
        (true && false) || (false && true) || !(false && false))

    x := 5
    x += 1
    if (x == 6){
        fmt.Println(x)
    } else {
        fmt.Println("You suck at code, bro.")
    }

    for i := 1; i <= 10; i++ {
        if i%2 == 0{
            fmt.Println(i, "even")
        } else {
            fmt.Println(i, "odd")
        }
    }

    for i :=1; i<= 100; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
    for i :=1; i<= 100; i++ {
        switch {
            case i%3 == 0 && i%5 ==0: fmt.Println("FizzBuzz")
            case i%3 == 0: fmt.Println("Fizz")
            case i%5 == 0: fmt.Println("Buzz")
        }
    }

    var q [5]int
    q[4] = 100
    fmt.Println(q)
    fmt.Println(len(q))
    for _, value := range q {
        fmt.Println(value)
    }
    fmt.Println()

    // Messing with arrays and slices and whatnot
    xp := [5]float64{
        98,
        93,
        77,
        82,
        83,
    }

    for _, value := range xp {
        fmt.Println(value)
    }
    xp_slc := xp[:3]
    fmt.Println(len(xp_slc))
    fmt.Println(5 == cap(xp_slc))
    fmt.Println(cap(xp_slc))
}
