package main

import "fmt"

//Write a function in Go that takes a slice of integers and returns a new slice containing only the unique integers from the original slice. The order of the elements in the returned slice should be the same as their first appearance in the original slice.

func UniqueEleSlice(s []int) []int {
	var uniqSlic []int

	for _, v := range s {
		flag := false

		for _, v1 := range uniqSlic {

			if v == v1 {
				flag = true
				break
			}

		}
		if !flag {
			uniqSlic = append(uniqSlic, v)

		}

	}

	return uniqSlic
}

func main() {
	mySlice := []int{2, 3, 4, 2, 42, 12, 21, 3, 2, 11, 4, 12}

	uniqSlic := UniqueEleSlice(mySlice)
	uniqSlic1 := UniqueEleSlice1(mySlice)

	fmt.Println(uniqSlic)
	fmt.Println(uniqSlic1)

}

//same program using maps
//because for the large slices it is not recomended to use 2 loops because of its higher time complexity
func UniqueEleSlice1(s []int) []int {
	var uniqueSlice []int
	uniqueMap := make(map[int]bool)
	for _, v := range s {
		exists := uniqueMap[v]
		//fmt.Println(uniqueMap)

		if !exists {
			uniqueMap[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}

	}
	return uniqueSlice
}
