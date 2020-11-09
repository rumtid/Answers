package leetcode

func getPermutation(n int, k int) string {
	set := make([]byte, n)
	for i := 0; i < n; i++ {
		set[i] = byte(i) + '1'
	}

	factorial := 1
	for i := 2; i < n; i++ {
		factorial *= i
	}

	k -= 1
	for s := set; n > 1; n-- {
		i := k / factorial
		if i != 0 {
			c := s[i]
			copy(s[1:i+1], s[0:i])
			s[0] = c
		}
		k -= factorial * i
		factorial /= n - 1
		s = s[1:]
	}

	return string(set)
}
