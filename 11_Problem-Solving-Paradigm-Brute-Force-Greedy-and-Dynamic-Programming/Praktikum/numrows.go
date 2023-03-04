package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func main() {
	numRows := 5
	res := generate(numRows)
	fmt.Println(res)
}
