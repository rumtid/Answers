package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var sets [][]int
	n := len(intervals)
	for i := 0; i < n; i++ {
		lo, hi := intervals[i][0], intervals[i][1]
		for i < n-1 && hi >= intervals[i+1][0] {
			if i++; intervals[i][1] > hi {
				hi = intervals[i][1]
			}
		}
		sets = append(sets, []int{lo, hi})
	}
	return sets
}
