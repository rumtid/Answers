package main

import "strconv"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	a, b, c := 1, 1, 0
	for i := 1; i < len(s); i++ {
		if n, _ := strconv.Atoi(s[i-1 : i+1]); n >= 10 && n <= 26 {
			c += a
		}
		if s[i] != '0' {
			c += b
		}
		a, b, c = b, c, 0
	}
	return b
}
