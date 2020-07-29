package main

import "math"

func divide(dividend int, divisor int) int {
	s, i := (dividend < 0) == (divisor < 0), 0
	x, y := abs(dividend), abs(divisor)
	for x >= y<<i {
		i++
	}
	var z int64
	for ; i >= 0; i-- {
		z <<= 1
		if x >= y<<i {
			x -= y << i
			z += 1
		}
	}
	if !s {
		z = -z
	}
	if z > math.MaxInt32 {
		return math.MaxInt32
	} else {
		return int(z)
	}
}

func abs(n int) int64 {
	if n < 0 {
		return -int64(n)
	} else {
		return int64(n)
	}
}
