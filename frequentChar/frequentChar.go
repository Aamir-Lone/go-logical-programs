package main

import (
	"fmt"
	"strings"
	"unicode"
)

//Write a function in Go named FindMostFrequentChar that accepts a string as input and returns the character that appears most frequently in the string, along with its count. The function should ignore spaces and be case-insensitive. If there are multiple characters with the same highest frequency, return any one of them.

func FindMostFrequentChar(s string) (rune, int) {
	input := strings.ToLower(s)
	charCount := make(map[rune]int)
	for _, v := range input {

		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			charCount[v]++
		}
	}
	var maxChar rune
	maxCount := 0
	for char, count := range charCount {
		if count > maxCount {
			maxChar = char
			maxCount = count
		}

	}

	return maxChar, maxCount
}

func main() {
	input := "This is a test string"
	char, count := FindMostFrequentChar(input)
	fmt.Printf("Most frequent character: %c (appears %d times)\n", char, count)

}
