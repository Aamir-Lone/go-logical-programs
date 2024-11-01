package main

//Write a Go function named IsAnagram that takes two strings as input and returns true if the strings are anagrams of each other, and false otherwise. An anagram is a word or phrase formed by rearranging the letters of another, using all the original letters exactly once. The function should be case-insensitive and ignore spaces and punctuation.

import (
	"fmt"
	"strings"
	"unicode"
)

func IsAnagram(s1, s2 string) bool {

	input1 := strings.ToLower(s1)
	input2 := strings.ToLower(s2)

	charCount1 := make(map[rune]int)
	charCount2 := make(map[rune]int)

	for _, r := range input1 {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			charCount1[r]++
		}
	}
	for _, r := range input2 {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			charCount2[r]++
		}
	}
	if len(charCount1) != len(charCount2) {
		return false
	}

	for k, v := range charCount1 {
		fmt.Println("val of k and v is:", k, v)
		if charCount2[k] != v {
			return false
		}
	}

	return true

}
func main() {

	str1 := "listen"
	str2 := "silent"

	if IsAnagram(str1, str2) {
		fmt.Println("given strings are anagram")
	} else {
		fmt.Println("given strings are not anagram")
	}

}
