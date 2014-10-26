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

    // var memory [0xFF]uint8

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

    // For loops

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

    // Switch and FizzBuzz

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
    
    // Let's append to a slice
    xp_slc_2 := append(xp_slc, 26)
    fmt.Println(4 == len(xp_slc_2))
    fmt.Println("Original array,", xp_slc)

    // Maps
    dict := make(map[string]int)
    fmt.Println("length of map: ", len(dict))
    dict["key"] = 10
    fmt.Println(dict["key"])
    fmt.Println("length of map: ", len(dict))
    // delete key
    fmt.Println("Removing the 'key' item.")
    delete(dict, "key")
    fmt.Println("length of map: ", len(dict))
    // check to see if we can find "key"
    if door, ok := dict["key"]; ok{
        fmt.Println(door, ok)
    } else {
        fmt.Println("That key value pair doesn't exist, buster.")
    }

    // Array, Slice, Map problems

    // 4th element == [3]
    // length of that slice would be 3 with a cap of 9
    // x[2:5] = ["c", "d", "e"]

    // smallest number problem
    smallest := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97,9,17,
    }
    i := smallest[0] // the control int
    for _, value := range smallest {
        if i > value {
            i =  value
        }
    }
    fmt.Println(i)
    avg_prblm := []float64{55, 8786, 942}
    fmt.Println(average(avg_prblm))
    fmt.Println()
    fmt.Println("Working on function stacks")
    fmt.Println(f1())
    the, greeting := multiple_values()
    fmt.Println(the, greeting)
    fmt.Println(add(1, 2, 3))
    add_slc := []int{1, 2, 3}
    fmt.Println(add(add_slc...))
}

// Functions chapter
func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }

    return total / float64(len(xs))
}

func f1() string {
    return f2()
}

func f2() (r string) {
    r = "Whuddup!"
    return
}

// Returning multiple values
func multiple_values() (r, s string) {
    r, s = "Sup", "kid?"
    return
}

// Variadic functions
func add(args ...int) (total int) {
    total = 0
    for _, v := range args {
        total += v
    }
    return
}
