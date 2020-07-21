package main

// you can also use imports, for example:
import (
	// "fmt"
	"math"
	"sort"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	minDis := make([]int, 0)

	// i != i+1
	// no value should be present [i] > [i+1] || [i] < [i+1]
	lenA := len(A)
	for i := 0; i < lenA; i++ {

		if i > 1 {
			for k := i - 1; k < lenA; k-- {
				if A[i] != A[k] {
					if !isValueLiesBW(A[i], A[k], A) {
						// fmt.Printf("I:%d IV: %d == k:%d kV:%d\n", i, A[i], k, A[k])
						d := math.Abs(float64(i) - float64(k))
						minDis = append(minDis, int(d))
						break
					}
				}
			}
		}

		for j := i + 1; j < lenA; j++ {
			if A[i] != A[j] {
				if !isValueLiesBW(A[i], A[j], A) {
					// fmt.Printf("I:%d IV: %d == J:%d JV:%d\n", i, A[i], j, A[j])
					d := math.Abs(float64(i) - float64(j))
					minDis = append(minDis, int(d))
					break
				}
			}
		}
	}

	sort.Ints(minDis)

	// fmt.Println("minDis>",minDis)
	minDistance := -1
	if len(minDis) > 0 {
		minDistance = minDis[0]
	}

	return minDistance
}

func isValueLiesBW(a, b int, arr []int) (present bool) {
	var min, max int
	if a > b {
		min, max = b, a
	} else {
		min, max = a, b
	}

	for _, a := range arr {
		if a > min && a < max {
			present = true
			break
		}
	}

	return
}
