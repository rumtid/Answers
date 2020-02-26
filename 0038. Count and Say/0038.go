package main

func countAndSay(n int) string {
	s := []byte{'1'}
	for i := 1; i < n; i++ {
		t, c, n := []byte{}, 1, s[0]
		for j := 1; j < len(s); j++ {
			if s[j] != n {
				t = append(t, byte(c+'0'), n)
				c, n = 1, s[j]
			} else {
				c++
			}
		}
		t = append(t, byte(c+'0'), n)
		s = t
	}
	return string(s)
}
