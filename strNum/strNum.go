package main

import (
	"fmt"
	"sort"
	"unicode"
)

func sortString(input string) string {
	var letters, digits []rune

	// Separate letters and digits
	for _, r := range input {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		} else if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}

	// Sort letters and digits independently
	sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
	sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })

	// Concatenate letters and digits
	return string(digits) + string(letters)
}

func main() {
	input := "b53663gdgsa63sg6s"
	result := sortString(input)
	fmt.Println(result)
}
