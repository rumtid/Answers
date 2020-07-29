package main

import (
	"fmt"
	"sort"
)

func Example_case1() {
	ans := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	for _, set := range ans {
		sort.Strings(set)
		fmt.Println(set)
	}
	// Unordered output:
	// [ate eat tea]
	// [nat tan]
	// [bat]
}
