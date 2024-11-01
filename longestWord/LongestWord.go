package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FindLongestWord(s string) string {

	input := strings.ToLower(s)

	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)

	})
	var longestWord string
	maxlength := 0
	for _, word := range words {

		length := len(word)
		if length > maxlength {
			maxlength = length
			longestWord = word
		}

	}
	return longestWord
}

func main() {

	myStr := "hello my name is aamir lone, i am currently learning go"

	longestWord := FindLongestWord(myStr)
	fmt.Println("the longest word is :", longestWord)

}
