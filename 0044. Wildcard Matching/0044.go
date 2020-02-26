package main

func isMatch(s string, p string) bool {
	var t, q string
	for len(s) > 0 {
		if len(p) > 0 && (p[0] == '?' || p[0] == s[0]) {
			s, p = s[1:], p[1:]
		} else if len(p) > 0 && p[0] == '*' {
			t, p, q = s, p[1:], p[1:]
		} else if len(t) == 0 {
			return false
		} else {
			s, t, p = t[1:], t[1:], q
		}
	}
	for ; len(p) > 0 && p[0] == '*'; p = p[1:] {
	}
	return len(p) == 0
}
