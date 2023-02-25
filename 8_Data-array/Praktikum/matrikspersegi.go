package main

import "fmt"

func main() {

	leftToRight := 15
	rightToLeft := 17

	absoluteDifference := leftToRight - rightToLeft
	if absoluteDifference < 0 {
		absoluteDifference *= -1
	}

	fmt.Printf("The absolute difference between the two diagonals is %d\n", absoluteDifference)
}
