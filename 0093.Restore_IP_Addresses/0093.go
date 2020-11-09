package leetcode

import (
	"strconv"
	"strings"
)

var (
	addr  []string
	addrs []string
)

func restoreIpAddresses(s string) []string {
	addrs = []string{}
	search(s, 4)
	return addrs
}

func search(s string, n int) {
	if n == 0 && len(s) == 0 {
		addrs = append(addrs, strings.Join(addr, "."))
	}
	if n == 0 || len(s) == 0 {
		return
	}
	for i := 1; i <= 3 && i <= len(s); i++ {
		if d, _ := strconv.Atoi(s[:i]); d <= 255 {
			addr = append(addr, s[:i])
			search(s[i:], n-1)
			addr = addr[:len(addr)-1]
		}
		if i == 1 && s[0] == '0' {
			break
		}
	}
}
