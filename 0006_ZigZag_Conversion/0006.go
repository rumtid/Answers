package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var t []byte
	n := (numRows - 1) * 2

	for i := 0; i < len(s); i += n {
		t = append(t, s[i])
	}
	for i := 1; i < numRows-1; i++ {
		for j := 0; ; j++ {
			if k := n*j + i; k < len(s) {
				t = append(t, s[k])
			} else {
				break
			}
			if k := n*(j+1) - i; k < len(s) {
				t = append(t, s[k])
			} else {
				break
			}
		}
	}
	for i := 0; ; i++ {
		if k := n*i + n/2; k < len(s) {
			t = append(t, s[k])
		} else {
			break
		}
	}

	return string(t)
}
