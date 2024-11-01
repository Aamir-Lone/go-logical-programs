package main

import "fmt"

func FilterEvenNumbers(n []int) []int {
	var evenNums []int
	for _, v := range n {
		if v%2 == 0 {

			evenNums = append(evenNums, v)
		}
	}

	return evenNums
}

func main() {

	nums := []int{2, 3, 4, 45, 7, 865, 54, 223, 1, 8, 6}
	evenNums := FilterEvenNumbers(nums)
	fmt.Println(evenNums)
}
