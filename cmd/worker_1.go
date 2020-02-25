package main

import "fmt"

func init() {
	Register("fibo", &workerFibo{})
}

type workerFibo struct{}

// Fibo 1,2,3,5,8
func (w *workerFibo) Do(limit int) {
	var first, last int
	current := 0
	first = 1
	for {
		current = first + last
		first = last
		last = current
		fmt.Println(current)

		if current > limit {
			break
		}
	}
}
