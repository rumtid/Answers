package main

import "fmt"

func Example_case1() {
	nums := []int{3, 2, 2, 3}
	len := removeElement(nums, 3)
	for i := 0; i < len; i++ {
		fmt.Println(nums[i])
	}
	// Unordered output:
	// 2
	// 2
}

func Example_case2() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	len := removeElement(nums, 2)
	for i := 0; i < len; i++ {
		fmt.Println(nums[i])
	}
	// Unordered output:
	// 0
	// 1
	// 3
	// 0
	// 4
}
