package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	if number < 1 {
		return -1
	}
	count := 0
	i := 2
	for {
		if isPrime(i) {
			count++
			if count == number {
				return i
			}
		}
		i++
	}
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
