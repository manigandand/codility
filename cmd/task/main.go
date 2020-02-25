package main

import (
	"fmt"
	"strings"
)

func main() {
	limit := 5

	out := make(chan int, limit)
	first := make(chan int)
	last := make(chan int)

	isLast := false
	for i := 0; i < limit; i++ {
		if i == 0 {
			go func() {
				first <- 1
				last <- 0
			}()
		}
		if i+1 == limit {
			isLast = true
		}

		f := <-first
		l := <-last

		// fmt.Printf("sending %d time islast: %v \n", i, isLast)
		go fibo(isLast, f, l, first, last, out)
	}

	// read all the response from the chan
	// out [ :close]
	for o := range out {
		print(o, " ")
	}

	fmt.Println()
	fmt.Println("I am done")
}

// should take first and last just prints the res
// output to chan wg *sync.WaitGroup,
func fibo(isLast bool, first, last int, firstC, lastC, out chan int) {
	current := first + last
	// fmt.Printf("INPUT  >>> F: %d, L: %d \nOUTPUT >>> F: %d, L: %d -> OUT: %d\n-----\n",
	// 	first, last, last, current, current)

	// write first, last and current
	out <- current
	if isLast {
		close(out)
	}

	firstC <- last
	lastC <- current
	if isLast {
		close(firstC)
		close(lastC)
	}
	return
}

func do(limit int) {
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

func test(first, last int, firstC, lastC chan int) {
	current := first + last
	fmt.Printf("OUTPUT >>> F: %d, L: %d -> OUT: %d\n",
		last, current, current)

	firstC <- last
	lastC <- current
	println(strings.Repeat("-", 15))
}
