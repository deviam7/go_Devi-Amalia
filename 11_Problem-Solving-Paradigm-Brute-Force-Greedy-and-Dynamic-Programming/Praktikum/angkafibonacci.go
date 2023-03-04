package main

import "fmt"

func fibo(n int) int {
	memo := make([]int, n+1)
	return fiboHelper(n, memo)
}

func fiboHelper(n int, memo []int) int {
	if n <= 1 {
		return n
	}
	if memo[n] > 0 {
		return memo[n]
	}
	memo[n] = fiboHelper(n-1, memo) + fiboHelper(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}
