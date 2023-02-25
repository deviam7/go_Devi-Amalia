package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	mergedArray := []string{}

	for _, value := range arrayA {
		if !contains(mergedArray, value) {
			mergedArray = append(mergedArray, value)
		}
	}

	for _, value := range arrayB {
		if !contains(mergedArray, value) {
			mergedArray = append(mergedArray, value)
		}
	}

	return mergedArray
}

// contains checks if a value is present in an array
func contains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "bryan"}))
	fmt.Println(ArrayMerge([]string{"devil jin", "sergei"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
