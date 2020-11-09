package leetcode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	buf := make([]bool, len(s2)+1)
	buf[0] = true
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i != 0 {
				buf[j] = buf[j] && s1[i-1] == s3[i+j-1]
			}
			if j != 0 {
				buf[j] = buf[j] || (buf[j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return buf[len(buf)-1]
}
