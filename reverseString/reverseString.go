package main

import (
	"fmt"
)

func ReverseString(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}

func main() {

	originalString := "aamirlone"

	reverseString := ReverseString(originalString)
	fmt.Println("Reversed string is", reverseString)
	fmt.Println("original string is", originalString)

}
