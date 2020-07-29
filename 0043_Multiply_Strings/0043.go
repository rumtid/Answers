package main

func multiply(num1 string, num2 string) string {
	var i, j int
	num3 := []byte{}
	nums := make([]int, 220)
	for i = 0; i < len(num1); i++ {
		for j = 0; j < len(num2); j++ {
			nums[i+j] += int((num1[len(num1)-i-1] - '0') * (num2[len(num2)-j-1] - '0'))
		}
	}
	for j = 219; j > 0 && nums[j] == 0; j-- {
	}
	for i, sum := 0, 0; i <= j || sum != 0; i++ {
		sum += nums[i]
		num3 = append(num3, byte(sum%10)+'0')
		sum /= 10
	}
	for i, j = 0, len(num3)-1; i < j; i, j = i+1, j-1 {
		num3[i], num3[j] = num3[j], num3[i]
	}
	return string(num3)
}
