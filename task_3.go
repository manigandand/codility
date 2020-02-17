package main

import "fmt"

func main() {
	a := []int{4, -1, 0, 3}
	b := []int{-2, 5, 0, 3}

	// a := []int{2, -2, -3, 3}
	// b := []int{0, 0, 4, -4}

	fmt.Println(Solution(a, b))
}

// Solution ...
func Solution(A []int, B []int) int {
	if len(A) != len(B) {
		return 0
	}
	// length N -> 1, N-1
	for index := 1; index <= len(A)-1; index++ {
		if findSliceSum(A, index) && findSliceSum(B, index) {
			return index
		}
	}

	return 0
}

func findSliceSum(a []int, k int) bool {
	slice1 := a[0:k]
	slice2 := a[k:]
	sum1 := sum(slice1)
	sum2 := sum(slice2)
	fmt.Println(a)
	fmt.Println(slice1, sum1)
	fmt.Println(slice2, sum2)
	fmt.Println("-----------------------------------------")

	return sum1 == sum2
}

func sum(a []int) (sum int) {
	for _, i := range a {
		sum += i
	}
	return
}
