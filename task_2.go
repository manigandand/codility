package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Solution(28))
}

// Solution ...
func Solution(N int) int {
	if N == 0 {
		return 0
	}
	inputSum := getSum(getInputStrList(N))
	return findNextMatchSum(N+1, inputSum)
}

func findNextMatchSum(start, sum int) int {
	for {
		if getSum(getInputStrList(start)) == sum {
			return start
		}
		// inc
		start++
	}
}

func getInputStrList(n int) []string {
	inputStr := strconv.Itoa(n)
	return strings.Split(inputStr, "")
}

func getSum(n []string) (sum int) {
	for _, i := range n {
		if v, err := strconv.Atoi(i); err == nil {
			sum += v
		}
	}
	return
}
