package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// Cari panjang string terpendek
	var n int
	if len(a) < len(b) {
		n = len(a)
	} else {
		n = len(b)
	}

	// Cari substring yang sama
	var result string
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			result += string(a[i])
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
