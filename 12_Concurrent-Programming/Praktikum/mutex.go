package main

import (
    "fmt"
    "sync"
)

var mutex = &sync.Mutex{}
var result = 1

func factorial(n int, wg *sync.WaitGroup) {
    defer wg.Done()

    mutex.Lock()
    for i := 1; i <= n; i++ {
        result *= i
    }
    mutex.Unlock()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go factorial(5, &wg)
    go factorial(6, &wg)

    wg.Wait()

    fmt.Println("Result:", result)
}