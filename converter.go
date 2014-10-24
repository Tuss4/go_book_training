package main

import (
    "fmt"
)

func main() {
    var input float64
    fmt.Print("Enter a temperature: ")
    fmt.Scanf("%f", &input)
    output := (input - 32) * 5/9
    fmt.Printf("%v degrees celsius.\n", output)

    var feet float64
    fmt.Print("Enter a measurement: ")
    fmt.Scanf("%f", &feet)
    meters := feet * 0.3048
    fmt.Printf("%v meters.\n", meters)
}
