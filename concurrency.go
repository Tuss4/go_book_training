// Will make each chapter its own go program from now on.
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func pinger(c chan string) {
    for i := 0;; i++ {
        c <- "ping"
    }
}

func ponger(c chan string) {
    for i := 0;; i++ {
        c <- "pong"
    }
}

func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    // for i := 0; i < 10; i++ {
    //     go f(i)
    // }
    // go f(0)
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
}
