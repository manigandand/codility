package main

import (
	"fmt"
	"time"
)

func main() {
	// random num generator
	minInterval := 0
	maxInterval := 10 // -1 X
	ns := time.Now().Nanosecond()
	// 1. divide secods by maxinterval
	// max 0 - 59
	rand := minInterval
	if ns == maxInterval {
		fmt.Println(rand)
		return
	} else { // 0 - 9
		rand = ns % (maxInterval + 1)
		fmt.Println("here", rand)
	}

	// TODO:check if not crossed more than maxInt
	fmt.Println(rand)
}

func test() {
	ip1 := []int{1, 2, 3, 4, 5}
	ip2 := []int{3, 4, 5}

	m := make(map[int]bool, len(ip2))
	for _, i := range ip2 {
		m[i] = true
	}

	res := make([]int, 0)
	for _, i := range ip1 { // N
		if ok := m[i]; !ok {
			res = append(res, i)
		}
	}

	fmt.Println(res)
}
