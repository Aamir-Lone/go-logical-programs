package main

import "fmt"

//Write a function in Go that takes a string and returns true if the string is a palindrome, and false otherwise.

func Palindrome(s string) bool {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}

	}
	return true
}

func main() {

	myStr := "Aamirlone"

	isPalindrome := Palindrome(myStr)
	if !isPalindrome {
		fmt.Println("string is not palindrome")

	} else {
		fmt.Println("String is Palindrome")

	}

}
