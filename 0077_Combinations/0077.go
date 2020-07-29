package main

var set []int
var sets [][]int

func combine(n int, k int) [][]int {
	sets = [][]int{}
	search(1, n, k)
	return sets
}

func search(start, end, k int) {
	if k == 0 {
		sets = append(sets, append([]int{}, set...))
		return
	}
	for i := start; i < end-k+2; i++ {
		set = append(set, i)
		search(i+1, end, k-1)
		set = set[:len(set)-1]
	}
}
