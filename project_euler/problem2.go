package main

import "fmt"

func fibbonaci(stream chan int) {
    x, y := 0, 1
    for {
        stream <- x
        stream <- y
        x = x + y
        y = x + y
        fmt.Println("x is now", x, "y is now", y)
    }
}

func main() {
    stream := make(chan int)
    go fibbonaci(stream)
    sum := 0
    for fibb := range stream {
        if fibb > 4000000 {
            break
        }
        if fibb % 2 == 0 {
            sum += fibb
        }
    }
    fmt.Println(sum)
}
