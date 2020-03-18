package main

import "fmt"

func Example_case1() {
	ans := findSubstring(
		"barfoothefoobarman",
		[]string{"foo", "bar"},
	)
	for _, num := range ans {
		fmt.Println(num)
	}
	// Unordered output:
	// 0
	// 9
}

func Example_case2() {
	ans := findSubstring(
		"wordgoodgoodgoodbestword",
		[]string{"word", "good", "best", "word"},
	)
	for _, num := range ans {
		fmt.Println(num)
	}
	// Unordered output:
}
