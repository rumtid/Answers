package leetcode

import "math"

func reverse(x int) int {
	var s, t int64
	if x >= 0 {
		s = int64(x)
	} else {
		s = -int64(x)
	}
	for s != 0 {
		t *= 10
		t += s % 10
		s /= 10
	}
	if x < 0 {
		t = -t
	}
	if t <= math.MaxInt32 && t >= math.MinInt32 {
		return int(t)
	} else {
		return 0
	}
}
