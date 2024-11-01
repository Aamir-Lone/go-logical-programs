package main

//Write a function in Go that takes a slice of integers and returns the sum of all the even numbers in the slice.

import "fmt"

func EvenNumSum(s []int) int {
	sum := 0
	for _, v := range s {
		if v%2 == 0 {
			sum += v
		}
	}
	return sum

}

func main() {

	mySlice := []int{2, 3, 4, 5, 2, 5, 1, 6, 7, 14}

	fmt.Println("sum of even numbers is:", EvenNumSum(mySlice))

}
