package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var lo, hi int = 0, len(intervals) - 1
	for ; lo <= hi && intervals[lo][1] < newInterval[0]; lo++ {
	}
	for ; lo <= hi && intervals[hi][0] > newInterval[1]; hi-- {
	}
	if lo <= hi && intervals[lo][0] < newInterval[0] {
		newInterval[0] = intervals[lo][0]
	}
	if lo <= hi && intervals[hi][1] > newInterval[1] {
		newInterval[1] = intervals[hi][1]
	}
	var sets [][]int
	for i := 0; i < lo; i++ {
		sets = append(sets, intervals[i])
	}
	sets = append(sets, newInterval)
	for i := hi + 1; i < len(intervals); i++ {
		sets = append(sets, intervals[i])
	}
	return sets
}
