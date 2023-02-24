package main

import "fmt"

func Primenumber(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Primenumber(1000000007))
	fmt.Println(Primenumber(13))
	fmt.Println(Primenumber(17))
	fmt.Println(Primenumber(20))
	fmt.Println(Primenumber(35))

}
