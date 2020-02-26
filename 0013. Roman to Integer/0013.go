package main

func romanToInt(s string) int {
	var i, j, k, n int
	for i = 0; i < len(s); i++ {
		j = runeToInt(s[i])
		if i+1 < len(s) {
			k = runeToInt(s[i+1])
		} else {
			k = 0
		}
		if j < k {
			n -= j
		} else {
			n += j
		}
	}
	return n
}

func runeToInt(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
