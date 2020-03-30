package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m+n-1, m-1, n-1
	for j >= 0 && k >= 0 {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			i, j = i-1, j-1
		} else {
			nums1[i] = nums2[k]
			i, k = i-1, k-1
		}
	}
	for k >= 0 {
		nums1[i] = nums2[k]
		i, k = i-1, k-1
	}
}
