package main

import "fmt"

func Mapping(slice []string) map[string]int {
	m := make(map[string]int)
	for _, str := range slice {
		if _, ok := m[str]; ok {
			m[str]++
		} else {
			m[str] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}
