package main

import (
	"strings"
)

func main() {
	first := make(chan int)
	for i := 0; i < 5; i++ {
		if i == 0 {
			go func() {
				first <- 1
			}()
		}

		input := <-first
		println("input>  ", input)
		go add(input, first)

		// close
		if i == 5 {
			close(first)
		}
	}
}

func add(i int, first chan int) {
	println("output> ", i*2)
	first <- i * 2
	println(strings.Repeat("-", 15))
}

func ex1() {
	first := make(chan int)
	go func() {
		go func() {
			first <- 1
			first <- 2
			// close(first)
		}()
		go func() {
			first <- 3
			first <- 4
			close(first)
		}()
	}()

	for o := range first {
		println(o)
	}
}
