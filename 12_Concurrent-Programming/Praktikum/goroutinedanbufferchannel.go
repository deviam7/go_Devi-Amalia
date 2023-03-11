package main

import "fmt"

func printMultiples(n int, ch chan<- int) {
    for i := 1; i <= n; i++ {
        if i%3 == 0 {
            ch <- i
        }
    }
    close(ch)
}

func main() {
    ch := make(chan int, 5)
    go printMultiples(15, ch)

    for i := 0; i < 5; i++ {
        fmt.Println(<-ch)
    }
}