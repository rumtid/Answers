package main

func lengthOfLongestSubstring(s string) int {
	var lo, hi, max int
	for ; hi < len(s); hi++ {
		for i := hi - 1; i >= lo; i-- {
			if s[i] == s[hi] {
				lo = i + 1
				break
			}
		}
		if hi-lo+1 > max {
			max = hi - lo + 1
		}
	}
	return max
}
