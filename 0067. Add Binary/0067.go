package main

func addBinary(a string, b string) string {
	var s byte
	var c []byte
	m, n := len(a), len(b)
	for i := 0; i < m || i < n || s != 0; i++ {
		if i < m {
			s += a[m-1-i] - '0'
		}
		if i < n {
			s += b[n-1-i] - '0'
		}
		c = append(c, s%2+'0')
		s /= 2
	}
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	return string(c)
}
