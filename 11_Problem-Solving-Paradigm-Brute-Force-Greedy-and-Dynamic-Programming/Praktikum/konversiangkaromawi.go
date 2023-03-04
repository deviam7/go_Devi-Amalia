package main

import "fmt"

// Fungsi untuk mengkonversi angka normal menjadi angka Romawi
func intToRoman(num int) string {
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

    result := ""

    for i := 0; i < len(values); i++ {
        for num >= values[i] {
            result += symbols[i]
            num -= values[i]
        }
    }

    return result
}

func main() {
    fmt.Println(intToRoman(4)) // Output: IV
    fmt.Println(intToRoman(9)) // Output: IX
    fmt.Println(intToRoman(23)) // Output: XXIII
    fmt.Println(intToRoman(2021)) // Output: MMXXI
    fmt.Println(intToRoman(1646)) // Output: MDCXLVI
}
