package leetcode

func longestPalindrome(s string) string {
	var t string

	for lo, hi := 0, len(s)-1; lo < len(s) && hi-lo+1 > len(t); lo++ {
		i, j, k := lo, hi, lo
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				k = i + 1
			}
		}
		if (hi-lo+1)-(k-lo)*2 > len(t) {
			t = s[k : hi-(k-lo)+1]
		}
	}

	for lo, hi := 0, len(s)-1; hi >= 0 && hi-lo+1 > len(t); hi-- {
		i, j, k := lo, hi, lo
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				k = i + 1
			}
		}
		if (hi-lo+1)-(k-lo)*2 > len(t) {
			t = s[k : hi-(k-lo)+1]
		}
	}

	return t
}
