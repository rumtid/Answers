package leetcode

import "sort"

var s [][]int

func combinationSum2(candidates []int, target int) [][]int {
	s = [][]int{}
	sort.Ints(candidates)
	combination([]int{}, candidates, target)
	return s
}

func combination(visited, candidates []int, target int) {
	for i, n := range candidates {
		if target == n {
			s = append(s, append([]int{}, append(visited, n)...))
		}
		if target <= n {
			break
		}
		if i != 0 && candidates[i] == candidates[i-1] {
			continue
		}
		combination(append(visited, n), candidates[i+1:], target-n)
	}
}
