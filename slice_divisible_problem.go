package main

import (
	"fmt"
	"sort"
)

func main() {
	// output = [5, 2, 1, 3, 7, 9, 12, 9, 11, 12, 3, 1]
	input := []int{5, 2, 7, 3, 1, 9, 12, 11, 9, 12, 3, 1}
	n := 2

	res := make([]int, 0)
	startIndex := 0
	divisiableIndex := 0

	inputLen := len(input)
	// keep appending in to the res
	for index, val := range input {
		if val%n == 0 {
			// get the slice of elements between this index
			// and sort this values
			divisiableIndex = index
			innerSlice := input[startIndex:divisiableIndex]
			innerSlice = sortInnerSlice(innerSlice)
			for _, v := range innerSlice {
				res = append(res, v)
			}
			// append the current index
			res = append(res, val)
			// set start index
			startIndex = index + 1
		}
	}

	// if reach last, just append the last values to res
	if divisiableIndex != inputLen-1 {
		for _, v := range input[startIndex:] {
			res = append(res, v)
		}
	}

	fmt.Println("Res: > ", res)
}

func sortInnerSlice(innerSlice []int) []int {
	sort.Ints(innerSlice)
	return innerSlice
}

func simpleSoln() {
	// output = [5, 2, 1, 3, 7, 9, 12, 9, 11, 12, 1, 3, 10]
	input := []int{5, 2, 7, 3, 1, 9, 12, 11, 9, 12, 3, 1, 10}
	n := 2

	startIndex := 0
	divisiableIndex := 0

	// inputLen := len(input)
	// keep appending in to the res
	for index, val := range input {
		if val%n == 0 {
			// get the slice of elements between this index
			// and sort this values
			divisiableIndex = index
			innerSlice := input[startIndex:divisiableIndex]
			sort.Ints(innerSlice)
			// set start index
			startIndex = index + 1
		}
	}

	fmt.Println("Res: > ", input)
}
