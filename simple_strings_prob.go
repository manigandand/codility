package main

import (
	"fmt"
	"sort"
)

// Ignore character: x -> d
// Sorting: [“abc”, “bbbcca”,  “cdd”,  “baba”,  “ccc”]
// a | b c a | c | b a| c
// -> [“a”,  “a”,  “a”,  “b”,  “b”,  “b”,  “c”,  “c”, “c”]

func main() {
	ignore := "d"
	inputs := []string{"abc", "bbbcca", "cdd", "baba", "ccc"}

	res := make([]string, 0)
	for _, ipStr := range inputs {
		isVisted := make(map[string]bool)
		for _, s := range ipStr {
			if string(s) != ignore {
				// make sure 1 unique char from this main string ipStr
				if ok := isVisted[string(s)]; !ok {
					// add it to list
					res = append(res, string(s))
					isVisted[string(s)] = true
				}
			}
		}
	}

	// i have unsorted rune chars
	fmt.Println(res)
	sort.Strings(res)
	fmt.Println(res)
}
