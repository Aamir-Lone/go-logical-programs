package main

import "fmt"

//Write a function in Go that takes a slice of integers and an integer n, and returns a new slice containing only the elements from the original slice that are greater than n.

func GreaterThan(s []int, n int) []int {

	var newslice []int
	for _, v := range s {

		if v > n {
			newslice = append(newslice, v)

		}

	}

	return newslice
}

func main() {
	num := 10
	numSlice := []int{10, 2, 3, 4, 53, 3, 1, 2, 5, 56, 67, 3, 2, 77}

	gtSlice := GreaterThan(numSlice, num)
	fmt.Println(gtSlice)

}
