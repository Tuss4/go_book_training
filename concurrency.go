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

func pinger(c chan<- string) {
    for i := 0;; i++ {
        c <- "ping"
    }
}

func ponger(c chan<- string) {
    for i := 0;; i++ {
        c <- "pong"
    }
}

func urkle(c chan<- string) {
    for i := 0;; i++ {
        c <- "Did I do thaaaat?"
    }
}

func printer(c <- chan string) {
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
    // var c chan string = make(chan string)

    // go pinger(c)
    // go ponger(c)
    // go urkle(c)
    // go printer(c)

    c1 := make(chan string, 1)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()
    go func() {
        for {
            c2 <- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()
    go func() {
        for {
            select {
            case msg1 := <- c1:
                fmt.Println("Message 1", msg1)
            case msg2 := <- c2:
                fmt.Println("Message 2", msg2)
            case <-time.After(time.Second * 2):
                fmt.Println("timeout") // Custom sleep function
            // default:
            //     fmt.Println("nothing ready")
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}

// Problems

/*
If you point to chan type it can only be sent to "c chan <- string"
If you point to the var it can only receive "c <- chan string"
*/

/*
To create a buffered channel that would be asynchronous add an int
inside of the make chan declaration.
*/
