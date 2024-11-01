package main

import (
	"errors"
	"fmt"
)

// Write a Go function named FindMinMax that accepts a slice of integers and returns two values: the minimum and maximum numbers in the slice. If the slice is empty, the function should return an error.
func FindMinMax(s []int) (int, int, error) {
	if len(s) == 0 {
		return 0, 0, errors.New("empty slice")
	}
	min := s[0]
	max := s[0]

	for _, v := range s {
		if v < min {
			min = v

		}
		if v > max {
			max = v
		}
	}
	return min, max, nil

}

func main() {

	myslice := []int{8889, 2, 3, 4, 25, 3, 1, 4, 33, 46, 664}

	min, max, err := FindMinMax(myslice)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println("minimum number is ", min)

	fmt.Println("maximum number is", max)

}
