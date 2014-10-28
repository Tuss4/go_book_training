package main

import (
    "fmt"
    "math"
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
    add_2 := func(x, y int) int {
        return x + y
    }
    fmt.Println(add_2(1, 567))
    // function having accesss to function scope aka a closure
    i_val := 0
    increment := func() int {
        i_val ++
        return i_val
    }
    fmt.Println(increment())
    fmt.Println(increment())

    nextEven := makeEvenGenerator()
    fmt.Println(nextEven())
    fmt.Println(nextEven())
    fmt.Println(nextEven())

    fmt.Println(factorial(4))
    // defer second()
    // first()
    // fmt.Println()
    // defer func() {
    //     str := recover()
    //     fmt.Println(str)
    // }()
    // panic("PANIC")

    // halving function
    fmt.Println(halves(1))
    fmt.Println(halves(2))
    // variadic function
    num_list := []int{1, 2, 3, 756, 8898, 12, 475856869}
    fmt.Println(find_the_greatest(num_list...))

    //Odd Generator
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())
    fmt.Println(nextOdd())
    fmt.Println(fibo(2))
    fmt.Println(fibo(14))
    p_var := 5
    fmt.Println(p_var)
    zero(&p_var)
    fmt.Println(&p_var) // To see the actual memory location
    fmt.Println(p_var)
    my_name := "Tomjo"
    fmt.Println(my_name)
    my_name = "Tomjo Soptame"
    fmt.Println(my_name)

    my_name = "Tomjo"
    fmt.Println(my_name)

    new_name := setMyName(&my_name)
    fmt.Println(my_name)
    fmt.Println(*new_name)
    xPtr := new(int)
    one(xPtr)
    fmt.Println(*xPtr)
    x_2_sqr := 1.5
    fmt.Println(x_2_sqr)
    square(&x_2_sqr)
    fmt.Println(x_2_sqr) // Should be 2.25
    m := 1
    n := 2
    swap(&m, &n)
    fmt.Printf("m: %v, n: %v\n", m, n)

    cir := Circle{0, 0, 5}
    fmt.Println(cir.x, cir.y, cir.r)
    fmt.Println(cir.area())
    tj := Person{"Tomjo"}
    tj.Talk()
    num_17 := new(Android)
    num_17.Talk()
    num_17.Person.Name = "Son Gohan"
    num_17.Talk()
    rect := Rectangle{0, 0, 10, 10}
    fmt.Println(rect.area())
    fmt.Println(totalArea(&cir, &rect))
}

func setMyName(name *string) *string {
    *name = "Tomjo Soptame"
    return name
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

// Closure function
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

// Recursive functions
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x -1)
}

// Defer, Panic & Recover
func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2nd")
}

// Functions problems
func halves(x uint) (uint, bool) {
    result := x/2
    var even bool
    if x%2 == 0 {
        even = true
    } else {
        even = false
    }
    return result, even
}

func find_the_greatest(args ...int) (i int) {
    i = 0 
    for _, value := range args {
        if i < value {
            i = value
        }
    }
    return i
}

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func fibo(x uint) uint {
    if x == 0 {
        return 0
    }
    if x == 1 {
        return 1
    }

    return fibo(x - 1) + fibo(x - 2)
}

// Pointers 
func zero(x *int) {
    /*
    So, if want my function to modify the value of the arg passed to it,
    I have to pass through a pointer that will point to the args location in memory.
    */
    *x = 0
}

func one(x *int) {
    *x = 1
}

// Pointers problems

/*
You get the memory address of a variable by pre-pending an '&' to it.
That turns it into a pointer which will point directly to its memory location.
*/

/*
You assign a value to a pointer by prepending a '*' to the variable, and then putting the new
value to the right side of the '='.
*/

/*
You can create a new pointer by using the built-in function 'new()' with the var type as an arg.
ex: x := new(string)
*/

/*Square problem*/
func square(x *float64) {
    *x = *x * *x
}

/*Swap problem*/

func swap(x, y *int) {
    *x, *y = *y, *x
}

// Structs
type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

type Person struct {
    Name string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
    Person
    Model string
}

type Shape interface {
    area() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

// Structs and Interfaces problems

/*
A method is a function that has a receiver to a struct.
*/

/*
If you want to have a struct be a field inside of another struct.
*/
