package main

func longestPalindrome(s string) string {
	var longest string
	var i, j, p, q, t int

	for i, j = 0, len(s)-1; i < len(s) && j-i+1 > len(longest); i++ {
		p, q, t = i, j, i
		for p < q {
			if s[p] != s[q] {
				t = p + 1
			}
			p++
			q--
		}
		if j-t-t+i+1 > len(longest) {
			longest = s[t : j-t+i+1]
		}
	}

	for i, j = 0, len(s)-1; j >= 0 && j-i+1 > len(longest); j-- {
		p, q, t = i, j, i
		for p < q {
			if s[p] != s[q] {
				t = p + 1
			}
			p++
			q--
		}
		if j-t-t+i+1 > len(longest) {
			longest = s[t : j-t+i+1]
		}
	}

	return longest
}
