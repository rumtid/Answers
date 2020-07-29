package main

import "fmt"

func Example_case1() {
	ans := permuteUnique([]int{1, 1, 2})
	for _, set := range ans {
		fmt.Println(set)
	}
	// Unordered output:
	// [1 1 2]
	// [1 2 1]
	// [2 1 1]
}
