package main

import "fmt"

func Example_case1() {
	ans := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// Unordered output:
	// 0
	// 9
}

func Example_case2() {
	ans := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// Unordered output:
}

func Example_case3() {
	ans := findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
	// Unordered output:
	// 8
}
