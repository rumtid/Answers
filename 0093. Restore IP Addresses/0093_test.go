package main

import "fmt"

func Example_case1() {
	ans := restoreIpAddresses("25525511135")
	for _, addr := range ans {
		fmt.Println(addr)
	}
	// Unordered output:
	// 255.255.11.135
	// 255.255.111.35
}
