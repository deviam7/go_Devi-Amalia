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
    ch := make(chan int)
    go printMultiples(15, ch)

    for multiple := range ch {
        fmt.Println(multiple)
    }
}