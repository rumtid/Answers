package leetcode

func mySqrt(x int) int {
	a := x
	for a*a > x {
		a = (x/a + a) / 2
	}
	return a
}
