package main

import (
	"fmt"
	"sync"
)

func countFrequency(text string, result map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, char := range text {
		result[char]++
	}
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur elit, sed do elusmod tempor incididunt ut labore et dolore magna aliqua"
	result := make(map[rune]int)

	var wg sync.WaitGroup
	wg.Add(4)

	length := len(text)
	quarter := length / 4

	go countFrequency(text[0:quarter], result, &wg)
	go countFrequency(text[quarter:2*quarter], result, &wg)
	go countFrequency(text[2*quarter:3*quarter], result, &wg)
	go countFrequency(text[3*quarter:length], result, &wg)

	wg.Wait()

	for char, freq := range result {
		fmt.Printf("%c : %d\n", char, freq)
	}
}