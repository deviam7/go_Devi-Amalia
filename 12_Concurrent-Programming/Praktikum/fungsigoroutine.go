package main

import (
    "fmt"
    "time"
)

func printMultiples(x int) {
    for i := 1; i <= 10; i++ {
        fmt.Println(i * x)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go func() {
        printMultiples(2)
    }()
    time.Sleep(3 * time.Second)
}