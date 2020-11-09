package leetcode

import "math"

func maxProfit(prices []int) int {
	var max, min int64 = 0, math.MaxInt32
	for _, price := range prices {
		num := int64(price)
		if num-min > max {
			max = num - min
		}
		if num < min {
			min = num
		}
	}
	return int(max)
}
