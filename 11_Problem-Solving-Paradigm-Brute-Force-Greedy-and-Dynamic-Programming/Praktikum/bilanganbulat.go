package main

import (
	"fmt"
	"strconv"
)

func toBinary(num int) string {
	if num == 0 {
		return "0"
	}
	binary := ""
	for num > 0 {
		remainder := num % 2
		binary = strconv.Itoa(remainder) + binary
		num /= 2
	}
	return binary
}

func getAns(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		binary := toBinary(i)
		count := 0
		for j := 0; j < len(binary); j++ {
			if binary[j] == '1' {
				count++
			}
		}
		ans[i] = count
	}
	return ans
}

func main() {
	n := 2
	ans := getAns(n)
	fmt.Println(ans)
}
