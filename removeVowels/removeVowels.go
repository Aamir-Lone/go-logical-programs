package main

//Write a function in Go that takes a string as input and returns a new string with all the vowels removed. The function should handle both uppercase and lowercase vowels (i.e., 'a', 'e', 'i', 'o', 'u' and 'A', 'E', 'I', 'O', 'U').

import "fmt"

func RemoveVowels(s string) string {

	var novovl string
	vowels := []string{"A", "a", "E", "e", "I", "i", "O", "o", "U", "u"}

	for _, ele := range s {
		flag := false

		for _, vovl := range vowels {
			if string(ele) == vovl {
				flag = true
				break
			}
		}
		if !flag {
			novovl += string(ele)

		}

	}
	return novovl
}

func main() {

	myStr := "Aamirlone q w -aeiou"

	newStr := RemoveVowels(myStr)
	newStr1 := RemoveVowels1(myStr)

	fmt.Println(newStr)
	fmt.Println(newStr1)

}

//other approach using maps

func RemoveVowels1(s string) string {
	novovl := ""
	vowels := map[rune]bool{'A': true, 'a': true, 'E': true, 'e': true, 'I': true, 'i': true, 'O': true, 'o': true, 'U': true, 'u': true}

	for _, ele := range s {
		if !vowels[ele] {
			novovl += string(ele)
		}
	}
	return novovl
}
