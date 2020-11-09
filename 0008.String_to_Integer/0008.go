package leetcode

import "math"

func myAtoi(str string) int {
	i := 0
	sign := 1
	var num int64
	var valid bool

	for ; i < len(str) && str[i] == ' '; i++ {
	}

	if i < len(str) {
		if str[i] == '+' {
			sign = 1
			i++
		} else if str[i] == '-' {
			sign = -1
			i++
		}
	}

Loop:
	for ; i < len(str); i++ {
		switch str[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num *= 10
			num += int64(str[i]) - int64('0')
			valid = true
		default:
			break Loop
		}
		if sign < 0 && -num < math.MinInt32 {
			return math.MinInt32
		}
		if sign > 0 && num > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	if sign < 0 {
		num = -num
	}
	if valid {
		return int(num)
	} else {
		return 0
	}
}
