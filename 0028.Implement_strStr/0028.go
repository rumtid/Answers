package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	t := make([]int, len(needle))
	t[0] = -1
	i, j := 1, 0
	for i < len(needle) {
		if needle[i] == needle[j] {
			t[i] = t[j]
		} else {
			t[i] = j
			for j >= 0 && needle[i] != needle[j] {
				j = t[j]
			}
		}
		i++
		j++
	}

	i, j = 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			if j+1 == len(needle) {
				return i - j
			} else {
				i++
				j++
			}
		} else {
			if t[j] >= 0 {
				j = t[j]
			} else {
				i++
				j = 0
			}
		}
	}

	return -1
}
