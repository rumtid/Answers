package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	mi, nn := 0, (len(nums1)+len(nums2)+1)/2
	lo, hi := 0, len(nums1)
	var a1, b1, a2, b2 int

	for {
		mi = (lo + hi) / 2
		if mi == 0 {
			a1 = math.MinInt32
		} else {
			a1 = nums1[mi-1]
		}
		if mi == len(nums1) {
			b1 = math.MaxInt32
		} else {
			b1 = nums1[mi]
		}
		if nn-mi == 0 {
			a2 = math.MinInt32
		} else {
			a2 = nums2[nn-mi-1]
		}
		if nn-mi == len(nums2) {
			b2 = math.MaxInt32
		} else {
			b2 = nums2[nn-mi]
		}

		if a1 > b2 {
			hi = mi - 1
			continue
		}
		if a2 > b1 {
			lo = mi + 1
			continue
		}

		if (len(nums1)+len(nums2))%2 == 0 {
			return float64(max(a1, a2)+min(b1, b2)) / 2
		} else {
			return float64(max(a1, a2))
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
