package main

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	switch i, j := len(s)-1, len(p)-1; {
	case p[j] >= 'a' && p[j] <= 'z':
		if i < 0 || s[i] != p[j] {
			return false
		}
		return isMatch(s[:i], p[:j])
	case p[j] == '.':
		if i < 0 {
			return false
		}
		return isMatch(s[:i], p[:j])
	case p[j] == '*':
		for c := p[j-1]; ; i-- {
			if isMatch(s[:i+1], p[:j-1]) {
				return true
			}
			if i < 0 || (c != '.' && s[i] != c) {
				return false
			}
		}
	}
	return false
}
