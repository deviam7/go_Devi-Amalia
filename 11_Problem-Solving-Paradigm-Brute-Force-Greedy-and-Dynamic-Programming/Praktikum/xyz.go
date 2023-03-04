package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := -10000; x <= 10000; x++ {
		for y := -10000; y <= 10000; y++ {
			z := a - x - y
			if x != y && x != z && y != z && x*y*z == b && x*x+y*y+z*z == c {
				fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("No solution")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
