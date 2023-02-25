package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	var result []int
	for _, r := range angka {
		str := string(r)
		digit, _ := strconv.Atoi(str)
		if len(strings.Replace(angka, str, "", -1)) == len(angka)-1 {
			result = append(result, digit)
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6,3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8,7,2,5,4]
}
