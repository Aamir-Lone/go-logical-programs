package main

//Write a function in Go that takes a slice of integers and an integer target, and returns true if any two numbers in the slice add up to the target, and false otherwise.

import "fmt"

func SumOfTwo(s []int, n int) bool {
	for i, v := range s {

		for j := i + 1; j < len(s); j++ {
			if (v + s[j]) == n {
				return true
			}
		}
	}
	return false
}

func main() {
	mySlice := []int{2, 3, 4, 1, 4, 3, 5, 8}
	num := 13

	if SumOfTwo(mySlice, num) {
		fmt.Println("two numbers in slice found that sum upto", num)

	} else {
		fmt.Println("No two numbers in slice sum up to", num)

	}

	if SumOfTwo1(mySlice, num) {
		fmt.Println("two numbers in slice found that sum upto", num)

	} else {
		fmt.Println("No two numbers in slice sum up to", num)

	}

}

//optimized approach using maps

func SumOfTwo1(s []int, target int) bool {
	numMap := make(map[int]bool)
	for _, num := range s {

		complement := target - num

		if numMap[complement] {
			return true
		}

		numMap[num] = true

	}
	return false

}
