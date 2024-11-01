package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountWordFrequencies(s string) map[string]int {
	freqMap := make(map[string]int)
	input := strings.ToLower(s)
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)

	})
	for _, v := range words {

		freqMap[v]++

	}

	return freqMap
}

func main() {

	input := "hello, hello,hello my name is Aamir Lone."
	freq := CountWordFrequencies(input)
	fmt.Println(freq)

}
