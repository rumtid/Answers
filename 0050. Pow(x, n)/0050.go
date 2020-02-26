package main

func myPow(x float64, n int) float64 {
	var m uint
	if n < 0 {
		m = uint(-int64(n))
		x = 1 / x
	} else {
		m = uint(n)
	}

	var p float64 = 1
	for m != 0 {
		if m%2 == 1 {
			p *= x
		}
		x *= x
		m /= 2
	}

	return p
}
