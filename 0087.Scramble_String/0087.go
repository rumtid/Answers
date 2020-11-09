package leetcode

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) == 1 {
		return false
	}
	s3 := []byte(s2)
	for i, j := 0, len(s3)-1; i < j; i, j = i+1, j-1 {
		s3[i], s3[j] = s3[j], s3[i]
	}
	return search(s1, s2) || search(s1, string(s3))
}

func search(s1, s2 string) bool {
	buf, num := [128]int{}, 0
	for i := 0; i < len(s1)-1; i++ {
		if buf[s1[i]]++; buf[s1[i]] > 0 {
			num++
		}
		if buf[s2[i]]--; buf[s2[i]] >= 0 {
			num--
		}
		if num == 0 && isScramble(s1[:i+1], s2[:i+1]) && isScramble(s1[i+1:], s2[i+1:]) {
			return true
		}
	}
	return false
}
